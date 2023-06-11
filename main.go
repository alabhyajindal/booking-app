package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const totalTickets int = 50
	var availableTickets uint = 50

	fmt.Printf("Welcome to the booking app for the %v\n", conferenceName)
	fmt.Printf("Total tickets for the event are %v, with %v tickets remaining\n", totalTickets, availableTickets)

	var firstName string
	var email string
	var ticketsWanted uint

	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)
	fmt.Println("What is your email?")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want?")
	fmt.Scan(&ticketsWanted)

	fmt.Printf("Thank you %v for your interest in booking %v tickets for the %v. You will recieve a email on %v shortly with a confirmation.\n", firstName, ticketsWanted, conferenceName, email)

	availableTickets = availableTickets - ticketsWanted
	fmt.Printf("%v tickets remaining.", availableTickets)
}
