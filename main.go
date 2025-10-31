package main

import "fmt"

func main() {

	fmt.Println("Get Your ticket to attend")

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("welcome to %v new booking app!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your ticket to attend")

}
