package main

import (
	"fmt"
	"greetings/greetings"
	"log"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	message, err := greetings.Hello("John")

	if err != nil {

		log.Fatalf("Unable to print greeting %v", err)
	}

	fmt.Println(message)
}
