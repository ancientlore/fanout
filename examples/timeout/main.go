package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		t := time.NewTimer(100 * time.Millisecond)
		select {
		case s := <-c:
			fmt.Println(s)
		case <-t.C: // HL
			fmt.Println("Timed out")
		}
	}()
	time.Sleep(150 * time.Millisecond)
}
