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

func main() {
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
