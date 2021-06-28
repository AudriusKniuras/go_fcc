package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func goroutines() {
	// goroutines create green (or virtual) threads
	// it's a lightway abstraction of OS threads, allows creating/destroying thousands of threads
	go sayHello() // without sleep goroutines() exits before sayHello() can execute
	time.Sleep(100 * time.Millisecond)

	// example of a race condition
	// race conditions can be checked with "go run -race ."
	msg := "Hello"
	go func() {
		fmt.Printf("%v - race condition\n", msg)
	}()
	msg = "Goodbye" // main thread is not paused until time.sleep is hit, so msg becomes Goodbye
	time.Sleep(100 * time.Millisecond)

	// proper way to use a variable
	msg = "Hello"
	go func(msg string) {
		fmt.Printf("%v - no race condition\n", msg)
	}(msg) // copy of msg passed to func, decoupled msg in a goroutine from msg in the function
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)

	// proper way to use goroutines with waitgroups (instead of sleeping)
	msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Printf("%v - goroutine with WaitGroup\n", msg)
		wg.Done() // remove 1 from wg. When wg becomes 0 it knows it finished everything
	}(msg)
	wg.Wait()

	// sync.Mutex and sync.RWMutex are used to make sure only 1 goroutine is manipulating data at a time (not implemented here)

	// Go will use CPU threads equal to available cores. Check and change with runtime.GOMAXPROCS
	// More threads may increase performance, need to check application with different GOMAXPROCS values
	fmt.Printf("%v - GOMAXPROCS default value", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Printf("Hello - goroutine#1\n")
}
