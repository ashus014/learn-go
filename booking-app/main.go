package main

import (
	"fmt"
	"strings"
)

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

	// fmt.Println(conferenceName)

	// var bookings = [50] string {"Ashu", "Akash", "Anu"}
	// var bookings [50] string

	// using slice, a dynamic list in Go
	var bookings [] string

	for {
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

		if userTicket > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTicket)
			continue
		}
	
		remainingTickets = remainingTickets - userTicket
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)
	
		// fmt.Printf("The whole slice : %v\n", bookings)
		// fmt.Printf("The first value : %v\n", bookings[0])
		// fmt.Printf("Slice type : %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))
	
	
		fmt.Printf("Thankyou %v %v for booking %v tickets. You'll receive a confirmation email address at %v\n", firstName, lastName, userTicket, email)
	
		// fmt.Printf("User %v booked %v tickets\n", firstName, userTicket)
		// fmt.Println("Memory location of userName Variable", &firstName)
	
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First Names: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
}