package main

import "github.com/linonymous/GoLearn/producer-consumer/producer"

import "github.com/linonymous/GoLearn/producer-consumer/consumer"

func main() {
	done := make(chan int)
	buffer := make(chan int)
	go producer.Produce(99, buffer)
	go consumer.Consume(buffer, done)
	<-done
	return
}
