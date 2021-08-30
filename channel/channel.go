package main

import (
	"fmt"
)

func foo(message <-chan string) { // Recieve only
	fmt.Println("FOO::DATA IN CHANNEL: ", <-message)
}

func bar(message chan<- string) { // Send only
	message <- "bar to main message passing"
}

func fooBar(message chan string) { // Recieve Send both
	fmt.Println("FOOBAR::DATA IN CHANNEL: ", <-message)
	message <- "foobar to main message passing"
}

func main() {

	message := make(chan string)

	go foo(message)
	message <- "main to foo message passing"

	go bar(message)
	fmt.Println("MAIN::DATA IN CHANNEL: ", <-message)

	go fooBar(message)
	message <- "main to fooBar message passing"
	fmt.Println("MAIN::DATA IN CHANNEL: ", <-message)
	fmt.Println("Main ended...")
}
