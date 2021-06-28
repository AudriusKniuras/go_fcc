package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroups to wait for goroutines and channels to sync data between them
var wg2 = sync.WaitGroup{}

func channels() {
	// int - type of data that flows through the channel
	ch := make(chan int)
	wg2.Add(2)
	// receiving channel go routine
	go func() {
		i := <-ch
		fmt.Printf("%v - channel example #1\n", i)
		wg2.Done()
	}()
	// sending channel go routine
	go func() {
		ch <- 42
		wg2.Done()
	}()
	wg2.Wait()

	// buffered channel - in case we have more senders than receivers
	// e.g. data is received fast, but processed for a long time or vice versa
	ch = make(chan int, 50) // create channel with a buffer of 50 (can hold 50 ints)
	wg2.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Printf("%v - buffered channel\n", i)
		}
		wg2.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // if we don't close it, receiving channel will not know if data is finished and it will cause a deadlock
		wg2.Done()
	}(ch)
	wg2.Wait()
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

// struct with no fields requires 0 memory allocation
// this channel cannot send any data through except that something was sent or received
// signal-only channel
var doneCh = make(chan struct{})

func channels2() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} // send a closure to the channel
}

func logger() {
	for { // infinity loop
		select {
		// if we receive message from logCh, we log, if we receive from doneCh, means channel is closing
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}

	// for entry := range logCh {
	// 	fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	// }
}
