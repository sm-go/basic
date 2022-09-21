package main

import (
	"fmt"
)

func main() {

	var title string = "Go Reference"
	fmt.Println(title)

	var totalTicket uint = 50
	var remainingTicket uint = 50
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTicket uint
	var booking [50]string

	fmt.Println("Enter your first name ...")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name ...")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address ...")
	fmt.Scan(&email)
	fmt.Println("Enter how many tickets would you get ...")
	fmt.Scan(&userTicket)
	remainingTicket = totalTicket - userTicket
	booking[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for your booked %v tickets and you will received via %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%d ticktes are remain \n", remainingTicket)

	fmt.Printf("Whole array is : %v \n", booking)
	fmt.Printf("The first value of array is : %v \n", booking[0])
	fmt.Printf("Type of array is : %T \n", booking)
	fmt.Printf("Length of array is : %v \n", len(booking))
}
