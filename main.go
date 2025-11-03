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

	var booking = [50]string{"Nama","Nicole","John"}
	var var booking = [50]string{"Nama","Nicole","John"}

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", 
	firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaininng for %v\n", remainingTickets, conferenceName)

	// userName = "John Doe"
	// userTickets = 2
	// fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", userName, userTickets)

}
