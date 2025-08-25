package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = []string{}

func main() {

	greetUsers()
	// todo()
	var first int
	var second int
	var operator string
	var task string
	var counter uint = 1
	myList := ToDoList{}
	for {

		fmt.Println("Enter your first number:")
		fmt.Scan(&first)

		fmt.Println("Enter your last number:")
		fmt.Scan(&second)

		fmt.Println("Enter your operator:")
		fmt.Scan(&operator)

		result := calculater(first, second, operator)
		fmt.Printf("Result: %d\n", result)

		fmt.Println("Welcome to the To-Do List Application!")
		fmt.Scan(&task)

		myList.AddTask(task)

		fmt.Printf("You have added %d tasks.\n", counter)
		fmt.Scan(&counter)
		myList.CompleteTask(int(counter))
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(
			firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}

		// myList.AddTask("Learn Go basics")
		// myList.AddTask("Build a To-Do app")
		// myList.AddTask("Push project to GitHub")

		// myList.ShowTasks()

		// myList.CompleteTask(2)

		myList.ShowTasks()
	}
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(
	userTickets uint,
	firstName string,
	lastName string,
	email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
