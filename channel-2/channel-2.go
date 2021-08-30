package main

import (
	"fmt"
)

func foo(message <-chan string) {
	fmt.Println("FOO::DATA IN CHANNEL: ", <-message)
}

func bar(message <-chan string) {
	fmt.Println("BAR::DATA IN CHANNEL: ", <-message)
}

func main() {

	message := make(chan string)

	go foo(message)
	go bar(message)
	message <- "message one"
	message <- "message two"

	fmt.Println("Main ended...")
}
