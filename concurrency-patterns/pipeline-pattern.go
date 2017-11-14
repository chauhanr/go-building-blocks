package concurrency_patterns

import (
	rand2 "math/rand"
	"fmt"
)

func PipelineGenerationFunc(integers []int) []int {
	var results []int

	done := make(chan interface{})
	defer close(done)

	inStream := generator(done, integers...)
	pipeline := multiply(done, add(done, multiply(done, inStream, 2), 1), 2)

	for v := range pipeline{
		results = append(results,v)
	}
	return results
}

func generator(done <-chan interface{}, integers ...int) <-chan int{
    inStream := make(chan int)

	go func(){
		defer close(inStream)
		for _, i := range integers{
			select{
			    case <-done:
			    	return
				case inStream <- i:
			}
		}
	}()
	return inStream
}


func multiply(done <-chan interface{}, inStream <-chan int, multipler int) <-chan int{
	multiInStream := make(chan int)

	go func(){
		defer close(multiInStream)
		for i := range inStream{
			select{
			case <-done:
				return
			case multiInStream <- i*multipler:
			}
		}
	}()
	return multiInStream
}

func add(done <-chan interface{}, inStream <-chan int, additive int) <-chan int{
	addInStream := make(chan int)

	go func(){
		defer close(addInStream)
		for i := range inStream{
			select{
			case <-done:
				return
			case addInStream <- i+additive:
			}
		}
	}()
	return addInStream
}



func FunctionRepaterFunc(){

	done := make(chan interface{})
	defer close(done)
	rand := func() interface{} { return rand2.Int()}

	for num := range take(done, funcRepeater(done, rand), 10){
		fmt.Printf("%d\n", num)
	}
}

func take(done <-chan interface{}, valueStream <-chan interface{} , num int) <-chan interface{}{
	takeStream := make(chan interface{})
	go func(){
		defer close(takeStream)
		for i:=0; i<num; i++{
			select {
			case <-done:
				return
			case takeStream <- <- valueStream:
			}
		}
	}()
	return takeStream

}

func funcRepeater( done <-chan interface{}, fn func() interface{}) <-chan interface{}{
	valueStream := make(chan interface{})

	go func(){
		defer close(valueStream)
		for {
			select{
				case <-done:
					return
				case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

