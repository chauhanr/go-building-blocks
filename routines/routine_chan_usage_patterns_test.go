package routines

import (
	"fmt"
	"testing"
	"time"
)

var cityResult = []string{"New Delhi", "Colombo", "Washington", "Timphu", "Seoul"}

/**
The test shows how the unbuffered channel must be used when two go routine need to communicate
1. the sender must be finite and must run in its routine. if it is not run in its own routine it will block the channel
and will cause the go runtime to throw a panic.
2. similarly the receiver function must also run in its own routine otherwise it causes the main test routine to block
and the last line is not able to run causing a deadlock too.

the final step asks the main routine to sleep allows both the routines to perform their operations otherwise we will
not get all the results.
*/
func TestUnbufferedChannelPattern(t *testing.T) {
	ch := make(chan string)
	cities := CityStruct{}

	go sendData(ch, cityResult)
	go getData(ch, &cities)
	time.Sleep(1e9)

	if cities.cities == nil {
		t.Errorf("The routines did not get the appropriate time to run.")
	}

	for i, city := range cities.cities {
		if city != cityResult[i] {
			t.Errorf("Expected %s and received %s \n", cityResult[i], city)
		} else {
			t.Logf("Capital : %s\n", city)
		}
	}
}

/**
The async channel with buffer differs from the unbuffered channel above
w.r.t to a channel size. This makes the channel non blocking until the buffer
size is breached upon which the blocking starts so in the example below if the channel
size is greater or equal to than the cityResult variable the sendData function can be
called from the main routine as the channel does not block.

If the buf size >= 5 then the test works but as soon as buf < 5 then the program fails
with a panic as the main routine blocks and does not allow the later steps to work.
the receiver however will always have to be in its own routine as it blocks the channel.
*/
func TestAsyncChannelsWithBufferPattern(t *testing.T) {
	buf := 6
	ch := make(chan string, buf)
	cities := CityStruct{}

	/*
				the following line can run in the main routine until
				the channel size (buf) is >= len(cityResult)
		    as soon as this condition is violated the channel blocks and we get an error from go.
	*/
	sendData(ch, cityResult)
	go getData(ch, &cities)
	time.Sleep(1e9)

	for i, city := range cities.cities {
		if city != cityResult[i] {
			t.Errorf("Expected %s and received %s \n", cityResult[i], city)
		} else {
			t.Logf("Capital : %s\n", city)
		}
	}
}

var bigStruct = []struct {
	bigArray []int
	result   int
}{
	{[]int{4, 8, 6, 9, 10, 13, 19, 41, 59, 81}, 250},
	{[]int{4, 8, 6, 9, 10, 13}, 50},
}

/**
Pattern 3 - this pattern passes a channel to a goroutine that will be used
to return a value that can be used as output. The channel will be used after
the routine to get the value.
retrieving the value from the channel is a blocking operation as it waits on the
channel to return a value.

This is also called a semaphore pattern.
*/

func TestChannelAsOutputRoutine(t *testing.T) {
	for _, array := range bigStruct {
		ch := make(chan int)
		go sum(array.bigArray, ch)
		sum := <-ch // this is the blocking operation.
		if sum != array.result {
			t.Errorf("Expected result %d but got %d\n", array.result, sum)
		} else {
			t.Logf("The correct sum was calculated: %d\n", sum)
		}
	}
}

func TestSemaphorePattern(t *testing.T) {
	s := make(semaphore, 5)
	go Add(3, 5, s)
	//time.Sleep(1e9)
}

/**
This test will create a channel and then pass that channel to the next consumer
this is called a channel factory pattern.
*/
func TestChannelFactoryPattern(t *testing.T) {
	suck(pump(5))
	time.Sleep(1e9)
}

// Test method to test out the pipe and filter pattern
// thr problem with the pipe and filter pattern is that we have not learnt how to close channels the proper way.
// therefore the example will return with an error that the channels were idel or were blocked.
func TestPipeFilterPattern(t *testing.T) {
	primes := sieve()
	for {
		prime, ok := <-primes
		if !ok {
			fmt.Println("primes channel has been closed.")
			break
		}
		fmt.Println(prime)
	}
}
