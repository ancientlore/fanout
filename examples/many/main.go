package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("Hello, world! (from goroutine #%d)\n", i) // HL
		}()
	}
	time.Sleep(250 * time.Millisecond)
}
