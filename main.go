package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const totalTickets = 50
	var availableTickets = 50

	fmt.Printf("Welcome to the booking app for the %v\n", conferenceName)
	fmt.Printf("Total tickets for the event are %v, with %v tickets remaining\n", totalTickets, availableTickets)
}
