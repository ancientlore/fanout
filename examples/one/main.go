package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, world!")
	}()
	time.Sleep(100 * time.Millisecond)
}
