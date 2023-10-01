package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceName is %T type\n conferenceTickets is %T type\n remainingTickets is %T type\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets form here to attend")

	fmt.Println(conferenceName)

	var userName string
	var userTicket int

	userName = "ashu"
	userTicket = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTicket)
}