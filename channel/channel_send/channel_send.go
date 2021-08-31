package main

import "fmt"

func sendVal(c chan int) {
	fmt.Println("SendVal Running...")
	c <- 8
}

func main() {

	valueChan := make(chan int)
	defer close(valueChan)

	go sendVal(valueChan)
	value := <-valueChan
	fmt.Println(value)
	fmt.Println("End main")
}
