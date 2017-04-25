package routines

import "testing"

func TestMultiplexServer(t *testing.T) {
	adder, quit := startServer(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.reply = make(chan int)
		adder <- req
	}

	for i := N - 1; i >= 0; i-- {
		value := <-reqs[i].reply
		if value != N+2*i {
			t.Errorf("Failed : Expected %d but got %d ", N+2*i, value)
		} else {
			t.Logf("Request %d is ok!\n", i)
		}
	}
	quit <- true
}

func f(left, right chan int) {
	left <- 1 + <-right
}

// this test method does nto have file to test.
func TestGoRoutineChaining(t *testing.T) {
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < 1000; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	x := <-leftmost
	t.Logf("The leftmost value in chain: %d ", x)
}
