package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avalible\n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for name
	fmt.Println("\nPlease enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("\nPlease enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("\nPlease enter your email:")
	fmt.Scan(&email)

	fmt.Println("\nHow many tickets you want?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. They will be sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
