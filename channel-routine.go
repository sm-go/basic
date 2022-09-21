package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //make the channel
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gophperID := <-c //receive value from a channel
		fmt.Println("ghoper ", gophperID, " has finished sleeping.")
	}
}

func sleepGopher(id int, c chan int) { //declear a channel as an argument
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id //send a value back to main
}
