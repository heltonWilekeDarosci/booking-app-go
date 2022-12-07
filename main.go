package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("What's your first name? ")
	fmt.Scan(&firstName)
	fmt.Println("What's your last name? ")
	fmt.Scan(&lastName)
	fmt.Println("What's your email? ")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want to purchase? ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
}
