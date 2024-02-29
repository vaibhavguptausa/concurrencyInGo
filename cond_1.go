package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func subscribe(c *sync.Cond, fn func()) {
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
	for i := 0; i < 1000; i++ {
		c := sync.NewCond(&sync.Mutex{})
		button := Button{Clicked: c}
		wg := sync.WaitGroup{}

		subscribe := func(c *sync.Cond, fn func(), ready *sync.WaitGroup) {
			wg1 := sync.WaitGroup{}
			wg1.Add(1)
			go func() {
				c.L.Lock()
				defer c.L.Unlock()
				wg1.Done()
				fmt.Println("lock taken------")
				ready.Done()
				c.Wait()
				fn()
			}()
			fmt.Println("waiting in routine------")
			wg1.Wait()
		}
		wg.Add(3)
		ready := &sync.WaitGroup{}
		ready.Add(3)
		subscribe(button.Clicked, func() {
			fmt.Println("I am first")
			wg.Done()
		}, ready)
		subscribe(button.Clicked, func() {
			fmt.Println("I am second")
			wg.Done()
		}, ready)
		subscribe(button.Clicked, func() {
			fmt.Println("I am third")
			wg.Done()
		}, ready)
		ready.Wait()
		button.Clicked.Broadcast()
		wg.Wait()
		fmt.Println("all waiting ------")
	}
}
