///This is my First try for use This beautiful language Golang.
package main

import (
	"fmt"
	"strings"
)

var confName string = "Go conf"

const confTicket int = 50

var remainingTicket uint = 50
var bookings = []string{}

func main() {

	greetuser(confName, confTicket, remainingTicket)

	for {
		///This is where I should put collect ticket and user data's
		firstName, lastName, userEmail, ticketNumber := getUserData()

		///This is where you take the validate parameter's.
		isValidEmail, isValidName, isValidTicketNumber := userValidate(firstName, lastName, userEmail, ticketNumber, remainingTicket)

		if isValidEmail && isValidName && isValidTicketNumber {
			///booking result output in here
			bookingResult(ticketNumber, remainingTicket, confName, firstName, lastName, userEmail)

			firstNames := printFirstName(bookings, firstName, lastName)
			fmt.Printf("This is all put bookings: %v\n", firstNames)

			if remainingTicket == 0 {
				// end of the program
				fmt.Println("Our confrence has been booked.")
				break

			}
		} else {
			if !isValidName {
				fmt.Printf("First and Last name must be 2 letter or more.\n")
			}
			if !isValidEmail {
				fmt.Printf("Please import a valid E-mail.\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("This is invalid number of ticket please add a valid number.\n")
			}
		}
	}
}

func greetuser(confName string, confTicket int, remainingTicket uint) {
	fmt.Println("Hello my friend")
	fmt.Printf("welcome to our new %v for begginers\n", confName)
	fmt.Printf("we have total of %v and we have %v left to book.\n", confTicket, remainingTicket)
	fmt.Println("please reserve you'r ticket as soon as possible")
}

func printFirstName(bookings []string, firstName string, lastName string) []string {
	bookings = append(bookings, firstName+" "+lastName)
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func userValidate(firstName string, lastName string, userEmail string, ticketNumber uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := ticketNumber <= remainingTicket && ticketNumber > 0
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var ticketNumber uint

	fmt.Print("please Enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Print("please Enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Print("please add your E-mail: ")
	fmt.Scan(&userEmail)
	fmt.Print("How many ticket's you want? ")
	fmt.Scan(&ticketNumber)
	return firstName, lastName, userEmail, ticketNumber
}

func bookingResult(ticketNumber uint, remainingTicket uint, confName string, firstName string, lastName string, userEmail string) {

	remainingTicket = remainingTicket - ticketNumber

	fmt.Printf("Thank you %v %v for your %v ticket reserved, we will send your confirmation to the %v \n", firstName, lastName, ticketNumber, userEmail)
	fmt.Printf("%v remaining ticket in %v \n", remainingTicket, confName)
}
