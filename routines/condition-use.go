package routines

import (
	"sync"
	"time"
	"fmt"
)

/**
	Here we implement the queue implementation using the cond condition to wait.
*/

func QueueImplUsingCond(){

	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration){
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Printf("Removed Element from queue\n")
		c.L.Unlock()
		c.Signal()
	}

	for i:=0; i<=10; i++{
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Printf("Adding Element to the queue\n")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second)
		c.L.Unlock()
	}
}

type Button struct{
	Clicked *sync.Cond
}

func BroadCastwithCondFunc(){

    button := Button{ Clicked : sync.NewCond(&sync.Mutex{})}

    subcribe := func(c *sync.Cond, fn func()){
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func(){
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subcribe(button.Clicked, func(){
		fmt.Printf("Maximixing the window \n")
		clickRegistered.Done()
	})

	subcribe(button.Clicked, func(){
		fmt.Printf("Display the annyoying dialog \n")
		clickRegistered.Done()
	})

	subcribe(button.Clicked, func(){
		fmt.Printf("Mouse Clicked. \n")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
