package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func(c <-chan bool) {
		var ok, b bool
		ok = true
		for ok {
			b, ok = <-c // HL
			fmt.Println(b, ok)
		}
	}(c)
	c <- false // HL
	c <- true  // HL
	close(c)   // HL
	time.Sleep(50 * time.Millisecond)
}
