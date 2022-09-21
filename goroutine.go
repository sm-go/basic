package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepGopher()
	}
	time.Sleep(4 * time.Second)
	fmt.Println("finished... after waithing 4 second")
}

func sleepGopher() {
	time.Sleep(2 * time.Second)
	fmt.Println("...snore...after waiting 2 second")
}
