package grpc

import (
	_ "github.com/arthasyou/grpc-consul-go/consul" // very important
	"github.com/arthasyou/grpc-consul-go/service"
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
