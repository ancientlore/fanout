package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func(c <-chan bool) {
		for b := range c { // HL
			fmt.Println(b)
		}
		fmt.Println("Done.")
	}(c)
	c <- false
	c <- true
	close(c) // HL
	time.Sleep(50 * time.Millisecond)
}
