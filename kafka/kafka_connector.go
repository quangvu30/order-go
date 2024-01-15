package kafka

import (
	"fmt"
	"os"

	"github.com/IBM/sarama"
	"github.com/quangvu30/order-go/logger"
)

type ConsumerInfo struct {
	Signal    chan os.Signal
	Partition int32
	Consumer  sarama.Consumer
}

type KafkaConnector struct {
	Brokers   []string
	Config    sarama.Config
	Producers map[string]sarama.SyncProducer
	Consumers map[string][]ConsumerInfo
}

func NewKafkaConnector(brokers []string) *KafkaConnector {
	cfg := sarama.NewConfig()
	return &KafkaConnector{
		Brokers:   brokers,
		Config:    *cfg,
		Producers: make(map[string]sarama.SyncProducer),
		Consumers: make(map[string][]ConsumerInfo),
	}
}

func (c *KafkaConnector) NewProducer(topic string) (*sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(c.Brokers, &c.Config)
	if err != nil {
		logger.Log.Error("Error creating Kafka producer: ", err.Error())
		return nil, err
	}
	c.Producers[topic] = producer
	return &producer, nil
}

func (c *KafkaConnector) NewConsumer(topic string, partition int32, onMessage func(data []byte)) {
	consumer, err := sarama.NewConsumer(c.Brokers, &c.Config)
	if err != nil {
		logger.Log.Error("Error creating Kafka consumer: ", err.Error())
		return
	}
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		logger.Log.Error("Error creating partition consumer: ", err.Error())
		return
	}

	// Handle signals to shutdown consumer
	c.Consumers[topic] = append(c.Consumers[topic], ConsumerInfo{
		Signal:    make(chan os.Signal, 1),
		Partition: partition,
		Consumer:  consumer,
	})

	logger.Log.Info(fmt.Sprintf("Consumer %s started!", topic))
	// Start a goroutine to consume messages
	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				onMessage(msg.Value)

			case err := <-partitionConsumer.Errors():
				logger.Log.Error("Error: ", err.Error())

			case <-c.Consumers[topic][partition].Signal:
				if err := partitionConsumer.Close(); err != nil {
					logger.Log.Error("Error closing partition consumer: ", err.Error())
				} else {
					logger.Log.Info("Closing partition consumer gracefully...")
				}
				return
			}
		}
	}()
}

func (c *KafkaConnector) TerminateConsumer(topic string, partition int32) {
	close(c.Consumers[topic][partition].Signal)
	if err := c.Consumers[topic][partition].Consumer.Close(); err != nil {
		logger.Log.Error("Error closing Kafka consumer:", err.Error())
	}
	c.Consumers[topic] = append(c.Consumers[topic][:partition], c.Consumers[topic][partition+1:]...)
}
