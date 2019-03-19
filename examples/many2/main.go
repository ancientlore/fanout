package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("Hello, world! (from goroutine #%d)\n", id) // HL
		}(i)
	}
	time.Sleep(250 * time.Millisecond)
}
