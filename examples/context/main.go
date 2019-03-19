package main

import (
	"context"
	"fmt"
	"time"
)

func rcv(ctx context.Context, c <-chan string) {
	select {
	case s := <-c:
		fmt.Println(s)
	case <-ctx.Done(): // HL
		fmt.Println("Timed out.")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // HL
	c := make(chan string)
	go rcv(ctx, c)
	go rcv(ctx, c)
	time.Sleep(150 * time.Millisecond)
}
