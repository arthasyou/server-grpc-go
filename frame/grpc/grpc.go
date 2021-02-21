package grpc

import (
	_ "github.com/luobin998877/go_grpc_with_consul/consul" // very important
	"github.com/luobin998877/go_grpc_with_consul/service"
	"github.com/spf13/viper"
)

// Start grpc service
func Start() {
	service.Register(&handler{})
	consulAddr := viper.GetString("Consul.addr")
	port := viper.GetInt("Service.port")
	name := viper.GetString("Service.name")
	tags := viper.GetStringSlice("Service.tags")
	s, err := service.CreateService(consulAddr, port, name, tags)
	if err != nil {
		return
	}
	go s.Start()
}
