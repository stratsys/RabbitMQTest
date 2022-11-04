package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/stratsys/RabbitMQTest/consumer"
	"github.com/stratsys/RabbitMQTest/producer"
)

func main() {

	var input string
	flag.StringVar(&input, "mode", "consum", "run as consum or produce")
	flag.Parse()
	if input == "producer" {
		producer.InitProducer()
	} else {
		consumer.InitConsumer()
	}
	fmt.Println("Shutdowm")
	os.Exit(0)

}
