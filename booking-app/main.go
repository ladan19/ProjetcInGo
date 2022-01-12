package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking applicarion\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTicket int
	// ask user for their name and ticket

	userName = "Ladan"
	userTicket = 2

	// I'm using printf to do this
	fmt.Printf("User %v booked %v tickets\n", userName, userTicket)
}
