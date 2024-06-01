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

	names := []string{"Gladys", "Samantha", "Darrin"}
	salams, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, salam := range salams {
		fmt.Println(salam)
	}
}
