package tpatterns

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func isTagged(url string) bool {
	r, err := http.Head(url)
	if err != nil {
		log.Println(err)
	}
	return r.StatusCode == http.StatusOK
}

// Server struct represents the server we need to poll for response.
type Server struct {
	version string
	url     string
	period  time.Duration

	mu  sync.RWMutex
	yes bool
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	if s.yes {
		fmt.Fprintf(w, "Yes. Golang version: %s is released", s.version)
	} else {
		fmt.Fprintf(w, "No. Golang version: %s is not released", s.version)
	}
	s.mu.Unlock()
}

// NewServer returns a new instance of hte server type.
func NewServer(version, url string, period time.Duration) *Server {
	s := &Server{version: version, url: url, period: period}
	go s.poll()
	return s
}

func (s *Server) poll() {
	for !isTagged(s.url) {
		time.Sleep(s.period)
	}
	s.mu.Lock()
	s.yes = true
	s.mu.Unlock()
}

func Crasher() {
	fmt.Println("Going down in flames")
	os.Exit(1)
}
