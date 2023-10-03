package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
)

func main() {
	// 1. 初始化客户端
	config := api.DefaultConfig()
	config.Address = "192.168.0.158:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 2. 服务注册
	// reg := &api.AgentServiceRegistration{
	// 	ID:   "myServiceID",
	// 	Name: "myService",
	// 	Port: 8080,
	// 	Check: &api.AgentServiceCheck{
	// 		HTTP:     "http://localhost:8080/health",
	// 		Interval: "10s",
	// 	},
	// }

	// client.Agent().ServiceRegister(reg)
	// defer client.Agent().ServiceDeregister("myServiceID")

	// HTTP server for health checks
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	go http.ListenAndServe(":8080", nil)

	// 3. 服务发现
	services, _, err := client.Catalog().Service("myService", "", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, service := range services {
		fmt.Printf("Found service: %v\n", service.ServiceName)
	}

	// 4. 键值存储
	pair, _, err := client.KV().Get("some_key", nil)
	if err != nil {
		log.Fatal(err)
	}
	if pair != nil {
		fmt.Printf("Value of 'some_key' is %s\n", pair.Value)
	}

	// Keep running
	select {}
}
