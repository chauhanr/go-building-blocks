package routines

// Request object will be a request to the multiplexing server.
type Request struct {
	a, b  int
	reply chan int
}

type operation func(a, b int) int

func run(op operation, request *Request) {
	request.reply <- op(request.a, request.b)
}

// The quit channel is present to shut the server down. The service channel of request methods takes a request and
// performs the operation,
func server(op operation, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op operation) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)

	go server(op, service, quit)
	return service, quit
}
