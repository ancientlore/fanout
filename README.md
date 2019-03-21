# fanout

A demo showing how to process incoming requests in parallel.

Useful to learn about:

* Concurrent execution (goroutines)
* Synchronization and messaging (channels)
* Multi-way concurrent control (`select`)
* Completing or canceling gracefully (`context.Context`)


## Build and Run

    $ go get github.com/ancientlore/fanout
    $ fanout

## Presentation

To run the presentation, use the [present](https://godoc.org/golang.org/x/tools/present) tool.

Install present:

    $ go get golang.org/x/tools/cmd/present

Run the presentation:

    $ present -notes -play

Navigate to http://127.0.0.1:3999/fanout.slide to see the presentation. Use the notes view by pressing `N` to see presenter notes.

## Good reading

* https://blog.golang.org/pipelines
* https://blog.golang.org/context
* https://golang.org/pkg/context/
* https://blog.golang.org/concurrency-is-not-parallelism
* https://talks.golang.org/2012/concurrency.slide
* https://swtch.com/~rsc/thread/
