package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings = [50]string{}

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
	fmt.Println(&remainingTickets)

	// & before variable is a pointer to its memory address
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firtsName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	bookings[0] = firtsName + " " + lastName

	fmt.Printf("The whole array: %v \n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	fmt.Printf("Array length: %v \n", len(bookings))
	fmt.Printf("Array type: %T \n", bookings)

	// I'm using printf to do this
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmantion email at %v \n", firtsName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaing  for %v \n", remainingTickets, conferenceName)

}
