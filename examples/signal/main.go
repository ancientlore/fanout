package main

import (
	"fmt"
	"time"
)

func waitForIt(id int, done <-chan bool) {
	<-done // HL
	fmt.Printf("%d is done!\n", id)
}

func main() {
	done := make(chan bool)
	go waitForIt(1, done)
	go waitForIt(2, done)
	time.Sleep(50 * time.Millisecond)
	done <- true // HL
	time.Sleep(50 * time.Millisecond)
}
