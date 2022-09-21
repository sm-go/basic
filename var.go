package main

import "fmt"

func main() {

	var myName = "Smith"
	fmt.Println("My name is ", myName)

	var myNickName string = "Jack"
	fmt.Println("My nick name is ", myNickName)

	name := "Leo Smith"
	fmt.Println("My name is ", name)

	world1, world2, _ := "hello", "world", "!"
	fmt.Println(world1, world2)
}
