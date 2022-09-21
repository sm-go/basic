package main

import (
	"fmt"
)

func main() {
	// name := "smithde"
	// fmt.Println("Before if")
	// if name == "smith" {
	// 	fmt.Println("inside if")
	// }
	// fmt.Println("after if")

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("How old are you?")
	// scanner.Scan()

	age := 16
	if age >= 18 {
		fmt.Println("You can ride")

	}
	if age >= 14 || age < 18 {
		fmt.Println("You can ride with your parents")
	}
	if age < 14 {
		fmt.Println("you can not ride")
		fmt.Printf("Wait %d years", 18-age)
	}
}
