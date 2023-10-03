package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"
)

// Config is the configuration for the application
type Config struct {
	Server struct {
		HTTP struct {
			Addr    string        `yaml:"addr"`
			Timeout time.Duration `yaml:"timeout"`
		} `yaml:"http"`
		GRPC struct {
			Addr    string        `yaml:"addr"`
			Timeout time.Duration `yaml:"timeout"`
		} `yaml:"grpc"`
	} `yaml:"server"`
	// ... Add other configuration fields here
	// For simplicity, only the 'Server' field is shown.
}

func TestConsulCURD(t *testing.T) {
	// 初始化客户端
	config := api.DefaultConfig()
	config.Address = "192.168.0.158:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 使用字符串形式写入配置
	// 	yamlData := `
	// server:
	//     http:
	//         addr: "0.0.0.0:30001"
	//         timeout: "1s"
	//     grpc:
	//         addr: "0.0.0.0:30002"
	//         timeout: "1s"
	// `
	// p := &api.KVPair{Key: "config/data", Value: []byte(yamlData)}
	// _, err = client.KV().Put(p, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 使用结构体形式写入配置
	var yamlData Config
	yamlData.Server.HTTP.Addr = "0.0.0.0:30001"

	// 把结构体序列化成[]byte, 满足api.KVPair.Value的类型
	data, err := yaml.Marshal(yamlData)
	if err != nil {
		log.Fatal(err)
	}

	// 写入配置
	p := &api.KVPair{Key: "config/data", Value: data}
	_, err = client.KV().Put(p, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 删除某个Key
	// kv1, err := client.KV().Delete("config/data", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("kv1", kv1) // Outputs:

	// 读取配置
	kv, _, err := client.KV().Get("config/data", nil)
	if err != nil {
		log.Fatal(err)
	}
	if kv == nil {
		log.Fatal("Config not found!")
	}

	// 把[]byte反序列化成结构体
	var cfg Config
	err = yaml.Unmarshal(kv.Value, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 输出配置
	fmt.Println(cfg) // Outputs: 0.0.0.0:30001
}
