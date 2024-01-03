package main

import "fmt"

func main() {
	var conferenceName = "Go Conference of 2024"
	const availableTickets = 50
	var remainingTicket = 46

	fmt.Println("Welcome to " + conferenceName)
	fmt.Println("No of tickets remaining ", remainingTicket)
	fmt.Printf("We are being booked quickly %v/%v tickets available", remainingTicket, availableTickets)
}
