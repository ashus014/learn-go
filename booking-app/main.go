package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	// syntatial sugar
	// const conferenceTickets int = 50
	conferenceTickets := 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T type\n conferenceTickets is %T type\n remainingTickets is %T type\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets form here to attend")

	fmt.Println(conferenceName)

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	// asking for user input
	fmt.Print("Enter your first name :: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name :: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email name :: ")
	fmt.Scan(&email)

	fmt.Print("Enter no. of tickets :: ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket

	fmt.Printf("Thankyou %v %v for booking %v tickets. You'll receive a confirmation email address at %v\n", firstName, lastName, userTicket, email)

	// fmt.Printf("User %v booked %v tickets\n", firstName, userTicket)
	// fmt.Println("Memory location of userName Variable", &firstName)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}