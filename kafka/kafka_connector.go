package kafka

import (
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
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
		log.Fatalf("Error creating Kafka producer: %v", err)
		return nil, err
	}
	c.Producers[topic] = producer
	return &producer, nil
}

func (c *KafkaConnector) NewConsumer(topic string, partition int32, onMessage func(data []byte)) {
	consumer, err := sarama.NewConsumer(c.Brokers, &c.Config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
		return
	}
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
		return
	}

	// Handle signals to shutdown consumer
	c.Consumers[topic] = append(c.Consumers[topic], ConsumerInfo{
		Signal:    make(chan os.Signal, 1),
		Partition: partition,
		Consumer:  consumer,
	})

	fmt.Printf("Consumer %s started!", topic)

	// Start a goroutine to consume messages
	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				onMessage(msg.Value)

			case err := <-partitionConsumer.Errors():
				fmt.Printf("Error: %v\n", err)

			case <-c.Consumers[topic][partition].Signal:
				fmt.Println("Closing partition consumer gracefully...")
				partitionConsumer.Close()
				return
			}
		}
	}()
}

func (c *KafkaConnector) TerminateConsumer(topic string, partition int32) {
	close(c.Consumers[topic][partition].Signal)
	if err := c.Consumers[topic][partition].Consumer.Close(); err != nil {
		fmt.Println("Error closing Kafka consumer:", err)
	}
	c.Consumers[topic] = append(c.Consumers[topic][:partition], c.Consumers[topic][partition+1:]...)
}
