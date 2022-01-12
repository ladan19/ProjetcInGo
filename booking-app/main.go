package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T\n", conferenceName, conferenceTickets)

	fmt.Printf("Welcome to %v booking applicarion\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firtsName string
	var lastName string
	var email string
	var userTicket int
	// ask user for their name and ticket

	// memory location of this variable
	fmt.Println(&conferenceName)

	// & before variable is a pointer to its memory address
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firtsName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	// I'm using printf to do this
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmantion email at %v \n", firtsName, lastName, userTicket, email)
}
