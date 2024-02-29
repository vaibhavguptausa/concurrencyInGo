package main

import (
	"fmt"
	"sync"
)

// Example shows how Once will focus on how many times do is called instead of getting called with different function hence the result is always  ?
func main() {
	once := sync.Once{}
	counter := 0
	for i := 0; i < 100; i++ {
		wgIncrement := sync.WaitGroup{}
		wgDecrement := sync.WaitGroup{}
		wgIncrement.Add(1)
		wgDecrement.Add(1)
		go func() {
			once.Do(func() {
				counter++
			})
			wgIncrement.Done()
		}()
		wgIncrement.Wait()
		go func() {
			once.Do(func() {
				counter--
			})
			wgDecrement.Done()
		}()
		wgDecrement.Wait()

	}
	fmt.Println("counter value is ----", counter)
}
