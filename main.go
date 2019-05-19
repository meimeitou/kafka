package main

import (
	"flag"
	"fmt"
	"kafka/kafkalibs"
)

var types string

func init() {
	flag.StringVar(&types, "t", "mete", "kafka types")
}

func main() {
	flag.Parse()
	switch types {
	case "mete":
		kafkalibs.MetadataTest()
	case "produce":
		fmt.Println("produce")
		kafkalibs.ProducerTest()
	case "consumer":
		fmt.Println("consumer")
		kafkalibs.ConsumerTest()
	case "cluster":
		kafkalibs.ClusterKafkaConsumer()
	case "cluster-part":
		kafkalibs.ClusterPartationCumsumer()
	default:
		fmt.Println("nothing")
	}
}
