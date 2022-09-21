package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson() *person {
	p := person{}
	fmt.Println(&p)
	return &p
}

func main() {
	person := newPerson()
	person2 := newPerson()
	fmt.Println(&person, &person2)
	//		fmt.Println(person{"Bob", 20})
	//		fmt.Println(person{name: "Alice", age: 23})
	//	}
}
