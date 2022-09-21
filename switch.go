package main

import (
	"fmt"
)

func main() {
	ans := 1
	switch ans {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Ten")
	}

}
