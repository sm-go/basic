package main

import (
	"fmt"
)

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, " is clean")
		} else {
			fmt.Println(room.name, " is dirty")
		}
	}
}

//something

func main() {
	var arr [3]int
	arr[0] = 343
	arr[2] = 44
	fmt.Println(arr)

	//array sum
	arr1 := [4]int{3, 6, 7, 8}
	fmt.Println(arr1)

	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println(sum)

	//two dimentional array
	// arr2D := [2][3]int{{1, 2, 6}, {3, 4, 7}}

	//slice array
	var orgarr [5]int = [5]int{12, 22, 36, 41, 50}
	var newarr []int = orgarr[1:3]
	// fmt.Println(newarr)
	fmt.Println(orgarr[:cap(newarr)])

	// Room struct

	rooms := [...]Room{
		{name: "Office"},
		{name: "Operation"},
		{name: "Reception"},
		{name: "Kitchen"},
	}
	checkCleanliness(rooms)
	fmt.Println("Performing cleaning...")
	rooms[1].cleaned = true
	rooms[3].cleaned = true
}
