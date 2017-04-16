package routines

import "fmt"

func sendData(ch chan string, c []string) {
	for _, city := range c {
		ch <- city
	}
}

func getData(ch chan string, c *CityStruct) {
	var input string

	for {
		input = <-ch
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
	for v := range ch {
		fmt.Printf("%d", v)
	}
}
