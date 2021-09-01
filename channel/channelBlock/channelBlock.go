package main

import (
	"fmt"
	"time"
)

func forever(message chan int) {
	for {
		//fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
		select {
		case msg, ok := <-message:
			if ok {
				fmt.Println("Message recieved: ", msg)
			}
		default:
			fmt.Println("Waiting for message in the channel...")

		}
	}
}

func main() {
	block := make(chan bool)
	messageToPass := make(chan int)
	go forever(messageToPass)

	messageToPass <- 3
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	messageToPass <- 30
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	messageToPass <- 100
	messageToPass <- 200
	time.Sleep(time.Second)

	<-block // block forever
}
