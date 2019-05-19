## golang kafka 示例代码

### 记录

export GO111MODULE=on
export GOPROXY=https://goproxy.io


go mod init ~/go/src/kafka
go get ./...

kafka-docker.md 记录了本地搭建kafka的方法

## 测试
启动：

### kafka元数据
go run main.go -t mete

### 异步生产者
go run main.go -t produce

### 普通消费者
go run main.go -t consumer

### 带记录的消费者
go run main.go -t cluster

### 带记录的分区消费者
go run main.go -t cluster-part
