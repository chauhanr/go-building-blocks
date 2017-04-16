package routines

type Empty interface{}
type semaphore chan Empty

// block n semaphore
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release the semaphore
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// noe implementing a mutex.
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal(n int) {
	s.V(1)
}
