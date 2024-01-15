package loadbalancer

var (
	orderQueue []string
	accountMap map[string]uint8 // accountCode : serverId
)

type LoadBalancer struct {
}

func NewLoadBalancer() *LoadBalancer {
	orderQueue = []string{}
	accountMap = make(map[string]uint8)
	return &LoadBalancer{}
}

func (b *LoadBalancer) Input(orderId string, accountCode string) {
	orderQueue = append(orderQueue, orderId)
	accountMap[accountCode] = 4
}
