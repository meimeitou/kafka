## 本地kafka集群docker安装

docker 安装

wurstmeister/zookeeper

wurstmeister/kafka

git：

https://github.com/wurstmeister/kafka-docker/

docker-compose.yml

- 将 KAFKA_ADVERTISED_HOST_NAME 换成自己的主机地址

```yaml
version: '2' 
services:
  zookeeper:
    image: wurstmeister/zookeeper
    ports:
      - "2181:2181" 
  kafka:
    image: wurstmeister/kafka
    ports:
      - "9092"
    environment:
      KAFKA_ADVERTISED_HOST_NAME: 192.168.124.11
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

```



启动：
docker-compose up -d


进入容器
docker exec -it wurkafka_kafka_1 /bin/bash

创建topic
$KAFKA_HOME/bin/kafka-topics.sh --create --topic test --zookeeper kafka_zookeeper_1:2181 --replication-factor 1 -partitions  1


 查看topic
$KAFKA_HOME/bin/kafka-topics.sh --zookeeper kafka_zookeeper_1:2181  --describe --topic test





发布消息: （输入若干条消息后 按^C 退出发布）



$KAFKA_HOME/bin/kafka-console-producer.sh --topic=test --broker-list kafka_kafka_1:9092

接收消息：



$KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server kafka_kafka_1:9092 --from-beginning --topic test

