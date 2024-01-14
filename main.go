package main

import (
	"fmt"
	"strings"
	"time"
)

func greetUsers(applicationName string) {
	fmt.Printf("Welcome to the %v ticker booking\n", applicationName)
}

func takeUserInputs() (string, string, string, uint) {
	var userTickets uint
	var userName string
	var lastName string
	var email string
	fmt.Println("Enter your first name:")
	fmt.Scan(&userName)

	fmt.Println("Enter you last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter you user email:")
	fmt.Scan(&email)

	fmt.Println("Ther number of tickets need to be booked:")
	fmt.Scan(&userTickets)

	return userName, lastName, email, userTickets
}

func validateUserInput(userName string, email string, lastName string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidEmail := strings.Contains(email, "@")
	isValidUserName := len(userName) > 2 && len(lastName) > 2
	isValiduserTickets := userTickets > 0 && userTickets >= remainingTickets

	return isValidEmail, isValidUserName, isValiduserTickets
}

func sendUserTicket(userName string, userTickets uint) {
	size := 15
	time.Sleep(50 * time.Second)
	for i := 1; i <= size; i++ {
		fmt.Print("=")
	}

	fmt.Printf("%v Ticket has been booked for %v", userName, userTickets)

	for i := 1; i <= size; i++ {
		fmt.Print("=")
	}

}
func main() {
	var applicationName = "Go-Airline"
	// applicationName := "Go-Airline"
	const totalTicket = 500
	var remainingTickets uint = 500
	greetUsers(applicationName)
	// var bookings []string

	type Userdetails struct {
		firtName    string
		lastName    string
		email       string
		userTickets uint
	}

	// or
	bookings := make([]Userdetails, 0)
	for {
		userName, lastName, email, userTickets := takeUserInputs()
		// isValidEmail, isValidUserName, isValiduserTickets := validateUserInput(userName, email, lastName, userTickets, remainingTickets)
		var isValidRemainingTicket bool = remainingTickets >= userTickets
		// if !isValidEmail || isValidUserName {
		// 	fmt.Print("Invalid userName or email ")
		// 	break
		// }

		// if isValiduserTickets {
		// 	fmt.Print("Ticket exceeded the availablity ")
		// 	break
		// }
		if isValidRemainingTicket {
			fmt.Print("Ticket can be booked")
			remainingTickets = remainingTickets - userTickets

			var userData = Userdetails{
				firtName:    userName,
				lastName:    lastName,
				email:       email,
				userTickets: userTickets,
			}
			bookings = append(bookings, userData)

			usernames := []string{}

			for _, booking := range bookings {

				usernames = append(usernames, booking.firtName)
			}

			go sendUserTicket(userName, userTickets)

		} else {
			fmt.Print("Ticket got exceeded tha the avalablity")
		}

	}

	// or
	// for remainingTickets>0 && len(bookings)<50{
	// fmt.Print("Above condton")
	//  }
}
