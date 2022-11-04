package main

import (
	"flag"
	"fmt"
	"main/consumer"
	"main/producer"
	"os"
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
