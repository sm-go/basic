package main

import (
	"fmt"
)

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"banana": 6,
		"cake":   7,
	}
	fmt.Println(mp)
	fmt.Println(mp["apple"])
	mp["apple"] = 4909
	fmt.Println(mp["apple"])

	//making a map

	myMap := make(map[string]int)

	myMap["favourite number"] = 7

	myItemMap := map[string]int{
		"item one":   100,
		"item two":   200,
		"item three": 300,
	}

	fmt.Println(myItemMap)
	delete(myItemMap, "item two")
	fmt.Println(myItemMap)

	//find the data in a array
	key, value := myItemMap["item one"]

	if !value {
		fmt.Println("Item One not found")
	} else {
		fmt.Println(value, key)
	}

	//find total in map
	mytotalItems := 0
	for _, amount := range myItemMap {
		mytotalItems += amount
	}
	fmt.Println("there are ", mytotalItems, "on my maps")

}
