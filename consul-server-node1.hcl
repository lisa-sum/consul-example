# 节点名称, 用于集群中的唯一标识
node_name = "consul-server-158"

# 本机IP
bind_addr = "192.168.0.158"

# 期望的集群节点数
bootstrap_expect = 3

# 重试加入集群的节点
retry_join = ["192.168.152","192.168.155"]

# 是否为服务端, ture为服务端, false为客户端
server    = true

# 是否开启bootstrap模式, 单机模式自动选举为leader
bootstrap = false

# 数据存储目录
data_dir = "/tmp/consul"

# 是否使用UI界面
# https://developer.hashicorp.com/consul/docs/agent/config/config-files#ui-parameters
ui_config {
  enabled = true
  content_path = "/ui/"
}

# 指定代理在哪个数据中心运行的字符串值
datacenter = "dc1"

# 日志级别
log_level  = "INFO"

# 服务监听地址
addresses {
  http = "0.0.0.0"
}

connect {
  enabled = true
}

