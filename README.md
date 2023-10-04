# consul-example

## Install 

Divided into binary installation and image installation, you can choose any way to install according to your preferences

### Binary

1. click [hashicorp](https://developer.hashicorp.com/consul/downloads) to download your os consul package
2. unzip consul
> if you use linux os you can use this command to unzip consul
> Apt package manager: apt install unzip -y
> Yum package manager: yum install unzip -y
> Dnf package manager: dnf install unzip -y

```shell
unzip consul.zip
```

### Image package
this used hashicorp/consul image

docker-compose.yml
```yaml
services:
  # consul 注册中心、配置中心
  consul:
    image: hashicorp/consul:1.16.1
    container_name: consul-stand-alone
    ports:
      - 8500:8500
    command:
      - agent
      - -server
      - -ui
      - -node=server-1
      - -bootstrap-expect=1
      - -client=0.0.0.0
```

## Use

### Stand-alone

#### Binary
dev mode

```shell
consul agent -dev -ui
```

#### Docker

```shell
docker-compose up -d
```

### Cluster
An example: a cluster of 3 servers with UI enabled
check [consul-server-node1.hcl](./consul.d/consul-server-node1.hcl), [consul-server-node2.hcl](./consul.d/consul-server-node2.hcl), [consul-server-node3.hcl](./consul.d/consul-server-node3.hcl) for more details
你需要根据你的实际情况来修改`consul-server-node{1,2,3}.hcl`的文件内容, 文件的字段说明如下:
> 根据你的需要, 参阅[官方文档](https://developer.hashicorp.com/consul/docs/agent/config/config-files#ui-parameters)配置文件, 根据你的实际需求来添加修改

- node_name: 节点名称
- bind_addr: 绑定地址, 一般是本机IP
- bootstrap_expect: 集群中server节点的数量, 一般是3,5,7等奇数数量
- retry_join: 集群中server节点的地址, 除了本机的地址, 其他server节点的地址都需要填写
- datacenter: 数据中心名称, 默认是`dc1`
- data_dir: 数据存储目录
- log_file: 日志文件路径
- log_level: 日志级别, 一般是`INFO`
- ui_config: Web UI选项
  - enabled: 是否启用Web UI, 默认是`true`
  - content_path : Web UI的路径, 默认是`/ui/`
- addresses: http和dns的监听地址
- ports: http和dns的监听端口, 默认是`8500`


- nohup: no hang up
- -config-file: config file path
- & : run in background

server node1:
```shell
nohup consul agent -config-file=/home/data/consul/consul.d/consul-server-node1.hcl &
```

server node2:
```shell
nohup consul agent -config-file=/home/data/consul/consul.d/consul-server-node2.hcl &
```

server node3:
```shell
nohup consul agent -config-file=/home/data/consul/consul.d/consul-server-node3.hcl &
```

Or 你可能更喜欢命令行的方式:
server node1:
```shell
consul agent -server -bootstrap-expect 3 -data-dir /tmp/consul -node=node_158 -bind=192.168.0.158 -rejoin -config-dir=/home/data/consul/consul.d/ -client 0.0.0.0
```
server node2:
```shell
consul agent -server -bootstrap-expect 3 -data-dir /tmp/consul -node=node_152 -bind=192.168.0.152 -rejoin -config-dir=/home/data/consul/consul.d/ -client 0.0.0.0
```

server node3:
```shell
consul agent -server -bootstrap-expect 3 -data-dir /tmp/consul -node=node_155 -bind=192.168.0.155 -rejoin -config-dir=/home/data/consul/consul.d/ -client 0.0.0.0
```

### Docker
> TODO, 与上面的cluster的docker-compose.yml类似, 只是需要修改`command`的内容


### Client
#### Golang
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
