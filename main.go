package main

import "fmt"

func main() {
	var confName = "Go conf"
	const confTicket = 50
	var remainingTicket = 50
	fmt.Println("Hello my friends")
	fmt.Printf("welcome to our new %v for begginers\n", confName)
	fmt.Printf("we have total of %v and we have %v left to book.\n", confTicket, remainingTicket)
	fmt.Println("please reserve you'r ticket as soon as possible")
}
