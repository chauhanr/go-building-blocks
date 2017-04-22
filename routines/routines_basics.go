package routines

import "fmt"

// CityStruct is the structure passed around
type CityStruct struct {
	cities []string
}

func sendData(ch chan string, c []string) {
	for _, city := range c {
		ch <- city
	}
	close(ch)
}

func getData(ch chan string, c *CityStruct) {
	for {
		input, ok := <-ch
		if !ok {
			fmt.Println("send data channel is closed.")
			break
		}
		c.cities = append(c.cities, input)
	}
}

func sum(bigArray []int, outCh chan int) {
	result := 0
	if bigArray != nil {
		for _, value := range bigArray {
			result += value
		}
		outCh <- result
	}
	outCh <- result
}

// Add method adds the integers sent in. Need to look at the semaphore pattern in more detail.
func Add(x int, y int, s semaphore) {
	s.Lock()
	fmt.Printf("Addition : %d \n", (x + y))
	s.Unlock()
}

func pump(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Printf("%d", v)
		}
	}()
}

/*
	Creating functions to explain the pipes and filter patterns with go routines.
	this pattern can be done using the channel directionality feature of channels.
	in the pipe filter pattern a function will take two channels as input and read from
	one channel process the item and write to the next channel.
	e.g.
		send := make(chan int)
		receive := make(chan string)
		go processChannels(send, receive)

    func processChannels(in <-chan int, out chan<- string){
			 for inValue := range in {
			     // process in value
					 out <- result
		 }
	 }

	 the generateNumbers and filter functions are called from the test file - routine_chan_usage_patterns_test.go
*/

// generateNumber function will generate numbers between 2 and a 100
func generateNumbers() chan int {
	out := make(chan int)
	go func() {
		for i := 2; i < 100; i++ {
			out <- i
		}
		defer close(out)
	}()
	return out
}

//  filter method will filter out any non prime numbers.
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			i, ok := <-in
			if !ok {
				break
			}
			if i%prime != 0 {
				out <- i
			}
		}
		defer close(out)
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generateNumbers()
		for {
			prime, ok := <-ch
			if !ok {
				break
			}
			ch = filter(ch, prime)
			out <- prime
		}
		defer close(out)
	}()
	return out
}
