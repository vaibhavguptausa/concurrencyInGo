package main

import (
	"fmt"
	"sync"
)

/*
What happens if we add to the wait-group inside the goroutine ? If we do that
it is possible that we reach the end even before the goroutine is scheduled
(and hence no addition to the wait-group counter), this means there will be no
blocking. Where have we done this in our codebase ?
*/

func main() {
	standardFlow()
	//cashfreeFlow()
}

func standardFlow() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("waiting for the end to come 1...")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("waiting for the end to come 2...")
	}()

	wg.Wait()
	fmt.Println("end has come for all of us")
}

func cashfreeFlow() {
	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		fmt.Println("waiting for the end to come 1...")
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		fmt.Println("waiting for the end to come 2...")
	}()

	wg.Wait()
	fmt.Println("end has come for all of us")
}
