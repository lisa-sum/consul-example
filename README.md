# consul-example

## Install

1. click [hashicorp](https://developer.hashicorp.com/consul/downloads) to download your os consul package
2. unzip consul
3. create consul config file and data dir
    ```shell
    mkdir -p /home/data/consul/consul.d
    ```
4. config consul file
    ```shell
    vi /home/data/consul/consul.d/consul.hcl
   ```
   
this is a cluster example:
   ```
    # /home/data/consul/consul.d/consul.hcl
    data_dir = "/tmp/consul" # 本地数据存储目录
    server = true # 服务器模式
    bootstrap_expect = 3 # 集群中期望的服务器节点数量
    node_name = "node-158"  # 对于不同的节点，请分别命名
    bind_addr = "192.168.0.158"  # 每个节点的 IP 地址
    retry_join = ["192.168.0.158", "192.168.0.155", "192.168.0.152"] # 配置集群中的其他节点
    addresses { # 配置 HTTP 和 DNS 服务器监听的地址
      http = "0.0.0.0"
    }
   ```
5. run consul server
    ```shell
    nohup consul agent -config-file=/home/data/consul/consul.d/consul.hcl &
    ```
6. add consul to path
    ```shell
    export PATH=$PATH:/path/to/consul
    ```
   
## Use

### Golang
1. install package
    ```shell
    go mod tidy
    ```

2. run test
    ```shell
    go run consul_test.go
    ```
3. run main
    ```shell
    go run main.go
    ```
4. check consul
visited link http://localhost:8080/health to check consul health if you can see the network status code is 200 then consul is ok
