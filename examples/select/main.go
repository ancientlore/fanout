package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done.")
				return
			case s := <-c:
				fmt.Println(s)
			}
		}
	}()

	c <- "Hello"
	c <- "World"
	close(done)
	time.Sleep(50 * time.Millisecond)
}
