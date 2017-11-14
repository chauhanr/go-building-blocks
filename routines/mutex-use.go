package routines

import (
	"sync"
	"fmt"
)

func MutexUseFunction(){
	var  count int
	var lock sync.Mutex
	/**
		Mutex is used to lock a critical section on the program to make is accessible by one program only.
	*/

	increment := func(){
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func(){
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrement: %d\n", count)
	}

	var arithmetic sync.WaitGroup

	for i :=0; i<=5; i++{
		arithmetic.Add(1)
		go func(){
			defer arithmetic.Done()
			increment()
		}()
	}

	for j :=0; j<=5; j++{
		arithmetic.Add(1)
		go func(){
			defer arithmetic.Done()
			decrement()
			}()
	}
	arithmetic.Wait()

	fmt.Printf("Arithmetic Routine complete\n")
}

func RWMutexUseFunc(){

}