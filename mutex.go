package main

import (
	"fmt"
	"sync"
)

type ButtonMutex struct {
	Clicked *sync.Cond
}

func subscribeM(c *sync.Cond, fn func()) {
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		c.L.Lock()
		fmt.Println("lock taken------")
		c.Wait()
		fn()
		c.L.Unlock()
	}()
	wg1.Wait()
}

func mutexExample() {
	var count int
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	fmt.Println("Complete")
}

func noMutexExample() {
	var count int
	increment := func() {

		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {

		count--
		fmt.Printf("decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	fmt.Println("Complete")
}

func main() {
	// if our lock works what should be the final value of count ?
	// mutex guarantees that the end result will be 0 while no mutex does not
	//mutexExample()
	noMutexExample()
}

func mutexWithCond() {
	c := sync.NewCond(&sync.Mutex{})
	button := ButtonMutex{Clicked: c}
	wg := sync.WaitGroup{}

	subscribe := func(c *sync.Cond, fn func()) {
		wg1 := sync.WaitGroup{}
		wg1.Add(1)
		go func() {
			wg1.Done()
			c.L.Lock()
			fmt.Println("lock taken------")
			fn()
			c.L.Unlock()
		}()
		wg1.Wait()
	}
	wg.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("I am first")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("I am second")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("I am third")
		wg.Done()
	})
	button.Clicked.Broadcast()
	wg.Wait()
}
