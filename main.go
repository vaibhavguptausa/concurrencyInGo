package main

import (
	"fmt"
	"sync"
	"time"
)

// Includes all examples of concurrency in go
func main() {
	//go SacchaiPrinter()
	//SacchaiKaRakhwala()
	//rangeOverSlice()
	rangeOverSliceCorrectly()
}

func SacchaiPrinter() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Sacchai kabhi chupti nahi hai")
}

func SacchaiKaRakhwala() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		SacchaiPrinter()
	}()
	wg.Wait() // ---> join point
}

func rangeOverSlice() {
	wg := sync.WaitGroup{}
	for _, val := range []string{"pen", "pineapple", "apple"} {
		wg.Add(1)
		// we exit the loop even before the goroutines are scheduled, hence val holds the last value (go smartly does not drop the variable val and pushes it into heap so it can be used, since it know it is still being referenced, )
		go func() {
			defer wg.Done()
			fmt.Println(val)
		}()
		// what will happen if we add a sleep ?
		//time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}

func rangeOverSliceCorrectly() {
	wg := sync.WaitGroup{}
	for _, val := range []string{"pen", "pineapple", "apple", "song"} {
		wg.Add(1)
		// the order will still not be deterministic but we are sure each goroutine will print exactly what it was called with
		go func(val string) {
			defer wg.Done()
			fmt.Println(val)
		}(val)
	}
	wg.Wait()
}
