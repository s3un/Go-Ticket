package main

import "fmt"

func main()  {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings[50]string

	fmt.Printf("wellcome to %v booking application", conferenceName)
	fmt.Println("we have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	// var email string
	// var userName string
	fmt.Println("What is your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("What is your last Name: ")
	fmt.Scan(&lastName)

	fmt.Printf("welcome to codebase %v, have a great time", firstName)
	

}