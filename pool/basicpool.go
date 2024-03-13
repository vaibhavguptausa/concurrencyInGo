package main

import (
	"fmt"
	"sync"
)

// we called get 3 times but the object got created just twice, since we put an instance back into it
func main() {
	mypool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new object")
			return struct{}{}
		},
	}
	mypool.Get()
	instance := mypool.Get()
	mypool.Put(instance)
	mypool.Get()
}
