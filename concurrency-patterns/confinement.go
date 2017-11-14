package concurrency_patterns

import (
	"fmt"
	"sync"
	"bytes"
)

func ConfinementWithChannels(){
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func(){
			defer close(results)
			for i:=0; i<5; i++{
				results <- i
			}
		}()
		return results
	}

	channelConsumer := func(results <-chan int){
		for result := range results{
			fmt.Printf("Received %d\n", result)
		}
		fmt.Printf("Completed Channel Consumption")
	}
	results := chanOwner()
	channelConsumer(results)
}


func ConfinementWithChannel2(){

	printData := func(wg *sync.WaitGroup, data []byte){
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data{
			fmt.Fprintf(&buff,"%c",b)
		}
		fmt.Printf(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
