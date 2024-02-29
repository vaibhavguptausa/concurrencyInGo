package main

import (
	"fmt"
	"sync"
	"time"
)

// shows the functionality of cond
// additional exercise - find why removeFromQueue1 works but   removeFromQueue2 does not

func removeFromQueue1(delay time.Duration, c *sync.Cond, queue *[]interface{}) {
	time.Sleep(delay)
	c.L.Lock()
	*queue = (*queue)[1:]
	fmt.Println("removed from queue")
	c.L.Unlock()
	c.Signal()
}

func removeFromQueue2(delay time.Duration, c *sync.Cond, queue []interface{}) {
	time.Sleep(delay)
	c.L.Lock()
	queue = queue[1:]
	fmt.Println("removed from queue")
	c.L.Unlock()
	c.Signal()
}

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)
	//removeFromQueue := func(delay time.Duration) {
	//	time.Sleep(delay)
	//	c.L.Lock()
	//	queue = queue[1:]
	//	fmt.Println("removed from queue")
	//	c.L.Unlock()
	//	c.Signal()
	//}
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			fmt.Println("waitiiiiiiiiiiiiing-------------------")
			c.Wait()
		}
		queue = append(queue, struct{}{})
		fmt.Println("added to queue")
		go removeFromQueue1(1*time.Second, c, &queue)
		//go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
