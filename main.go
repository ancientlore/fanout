package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// S holds the data we produce and consume.
type S struct {
	Name string // Name of item
	ID   int    // Item ID
}

// producer generates up to count data items on the returned channel.
func producer(ctx context.Context, count int) <-chan S {
	// make result channel
	c := make(chan S)

	// start goroutine that produces results
	go func() {
		// done is closed if we are supposed to stop
		done := ctx.Done()

		// important - we need to close the result channel when finished
		defer close(c) // HL

		// produce data items on the channel
		for i := 0; i < count; i++ {
			s := S{Name: fmt.Sprintf("Item number %d", i+1), ID: i}
			select {
			case <-done: // HL
				// processing cancelled; return
				return
			case c <- s: // HL
				// new item sent
			}
		}
	}()

	// return the channel that will have results
	return c
}

// consumer spawns threadCount goroutines to processes messages on c,
// and waits for them to finish.
func consumer(ctx context.Context, c <-chan S, threadCount int) {
	// Initialize wait group with number of threads/goroutines
	var wg sync.WaitGroup
	wg.Add(threadCount)

	// start goroutines to process the channel
	for i := 0; i < threadCount; i++ {
		go func(id int) { // HL
			// process until cancelled or the channel is closed
			processor(ctx, id, c) // HL
			// decrement wait group counter
			wg.Done() // HL
		}(i)
	}

	// Wait until all threads are done
	wg.Wait() // HL
}

// processor reads messages on c and processes them.
func processor(ctx context.Context, id int, c <-chan S) {
	done := ctx.Done()
	for {
		select {
		case <-done: // HL
			// processing cancelled; return
			return
		case item, ok := <-c: // HL
			if !ok {
				// channel closed; all done
				return
			}
			// Here we do the time-consuming processing step, which in this case
			// is fast.
			fmt.Printf("%4d: Processing %4d %q\n", id, item.ID, item.Name)
		}
	}
}

var (
	itemCount  = flag.Int("count", 100, "Number of items to produce")
	goroutines = flag.Int("threads", 10, "Number of goroutines for processing")
)

func main() {
	// Parse flags
	flag.Parse()

	// Create context, enables us to cancel the process cleanly
	// See:
	// https://golang.org/pkg/context/
	// https://blog.golang.org/context
	// https://blog.golang.org/pipelines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle Ctrl-Break, Ctrl-C
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func(ctx context.Context) {
		done := ctx.Done()
		select {
		case x := <-sig:
			fmt.Println("Received", x)
			cancel()
			signal.Stop(sig)
		case <-done:
			fmt.Println("Done")
		}
	}(ctx)

	// Start the producer
	c := producer(ctx, *itemCount)

	// Run the consumer, which spawns goroutines
	consumer(ctx, c, *goroutines)
}
