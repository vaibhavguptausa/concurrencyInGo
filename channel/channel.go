package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func theChannel() {

	stream := make(chan string)
	go func() {
		stream <- "ganga putra ho to beh jao"
		stream <- "nahi ho tab bhi beh jao"
	}()

	// aayaKya here is a boolean which is an optional output, it just shows whether you are reading a message which was published by someone else
	// or you are reading a default value from a  closed channel
	kabootarKaMessage, aayaKya := <-stream
	fmt.Printf("%v, %v\n", aayaKya, kabootarKaMessage)

	// each of the reads give you one value, once you get the value you stop reading
	// if you want to read the second message you need to read again like mentioned below
	// our main goroutine will wait till both the reads are done and hence  anonymous go
	// func will always return before main go routine

	fmt.Println(<-stream)
}

func closedChannel() {
	stream1 := make(chan int)
	stream2 := make(chan int)

	go func() {
		defer close(stream2)
	}()
	go func() {
		defer close(stream1)
		stream1 <- 1
	}()

	val1, ok1 := <-stream1
	val2, ok2 := <-stream2
	fmt.Printf("%v, %v\n", ok1, val1) // shows the value was produced somewhere else and is not default by printing true
	fmt.Printf("%v, %v", ok2, val2)
}

// use this to iterate over the channel
func iteratingOverAChannel() {
	stream := make(chan int)
	go func() {
		defer close(stream)
		for i := 0; i < 10; i++ {

			stream <- i
			fmt.Println("pushed " + fmt.Sprintf("%d", i))
		}
	}()

	// the loop is reading from the channel until it is closed.
	for val := range stream {
		fmt.Println(val)
	}

}

// as soon as you close the channel all the blocked goroutines will resume.
func signalByClosing() {
	begin := make(chan interface{})
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("begun %d \n", i)
		}(i)
	}
	close(begin)
	wg.Wait()
}

// since we have a buffered channel here, we can push multiple values to it without it being read somewhere else.

func aBufferedChannel() {
	var stdOutBuff bytes.Buffer
	defer stdOutBuff.WriteTo(os.Stdout)

	initStream := make(chan int, 5)

	go func() {
		defer fmt.Fprintf(&stdOutBuff, "producer done.\n")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdOutBuff, "sending: %d \n", i)
			initStream <- i
		}
	}()
	for val := range initStream {
		fmt.Fprintf(&stdOutBuff, "received %d\n", val)
		if val == 4 {
			close(initStream)
		}
	}
}

// This method just shows that we have to close the channel or we will have deadlock, if we are closing the channel in the end we can iterate without any issue, as soon as we put a message in the channel we will receive it outside
func aBufferedChannel2() {
	var stdOutBuff bytes.Buffer
	defer stdOutBuff.WriteTo(os.Stdout)

	initStream := make(chan int, 5)

	go func() {
		defer close(initStream)
		for i := 0; i < 5; i++ {
			fmt.Println("sending ", i)
			time.Sleep(1 * time.Second)
			initStream <- i
		}
	}()
	for val := range initStream {
		fmt.Println("receiving  ", val)
		if val == 4 {
		}
	}
}

func main() {
	//theChannel()
	//closedChannel()
	//iteratingOverAChannel()
	//signalByClosing()
	//aBufferedChannel()
	aBufferedChannel2()
}
