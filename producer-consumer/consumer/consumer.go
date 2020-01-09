package consumer

import "fmt"

import "time"

func Consume(buffer <-chan int, done chan<- int) {
	for {
		i := <-buffer
		time.Sleep(2)
		fmt.Printf("Received %d\n", i)
		done <- 5
	}
}
