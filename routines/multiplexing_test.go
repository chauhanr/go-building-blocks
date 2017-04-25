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
