package nsq

import (
	"github.com/luobin998877/go_utility/nsq"
	"github.com/spf13/viper"
)

// Connect nsq node
func Connect() {
	addr := viper.GetString("Producer.addr")
	nsq.ProducerStart(addr)
}

// Publish message to nsq
func Publish(topic string, data []byte) {
	nsq.Publish(topic, data)
}

// Stop producer
func Stop() {
	nsq.ProducerStop()
}
