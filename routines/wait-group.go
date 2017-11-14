package routines

import (
	"sync"
	"fmt"
	"time"
)

func ForkJoinFunction (){

	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "welcome", "good day"}{
		wg.Add(1)
		go func(sal string){
			defer wg.Done()
			fmt.Printf("%s \n",sal)
		}(salutation)
	}
	wg.Wait()
}


func SleepingRoutines(){

	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Printf("1st routine sleeping\n")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Printf("2nd goroutine sleeping\n")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Printf("All routines complete .. \n")
}

func BulkWGUseage(){
	var wg sync.WaitGroup
	hello := func(wg *sync.WaitGroup, id int){
		defer wg.Done()
		fmt.Printf("Hello Process : %d\n", id)
	}
	const NumGreets=5
	wg.Add(NumGreets)
	for i:=0; i<NumGreets; i++{
		go hello(&wg, i)
	}
	wg.Wait()
}

