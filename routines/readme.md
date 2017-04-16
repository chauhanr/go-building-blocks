# Golang Routines

### Concurrency Parallelism and Go Routines

Concurrency is the model where the programming language allows for multiple threads to run which share the same memory space of the process they run in, but many a times the concurrency model does not allow for threads to efficiently run on multiple cores. There are various languages that provide concurrency but do not efficiently parallelize, parallelization comes into the picture when the language can effectively move these threads to processors.

Multi-threaded applications are notoriously difficult to get right as there is always shared data in memory. The solution to allow for concurrent access to memory in data is to use synchronization. This is however inefficient way of making concurrent programs as it leads to more errors and inefficiencies. Go has sync options but it does not prescribe the “thread per connection model” of programming.

Go on the other hand uses what is called CSP (communicating sequential process) or also called message passing model (Erlang). In Go the parts that run concurrently are called Go routines and there is not one-to-one mapping between goroutines and OS threads : goroutine are mapped onto (multiplexed) one or more threads, accordingly to their availability; this is accomplished by goroutine scheduler in the Go runtime.  Go highly discourages the use of sync package but use the channels in routines to sync the routines.


  Go routines are lightweight and much lighter than a thread. The small footprint of the routine is because it dose not make a fixed size stack for the thread or execution unit. Typically the  go routine uses 4K size and then grows unlike other languages where stack is set to 2 MB which can waste space when many small threads are spawned.


  As mentioned earlier the go routines are not scheduled by the OS but the go runtime. The routines are multiplexed over a set of threads on the OS. the go runtime is smart enough to realize when a system call blocks the routine so that it can switch the routine.


  Concurrency are of two styles
  1. deterministic (well ordered and understood) - go routine promotes this as channels and communication through them makes the program easier to understand.
  2. non deterministic (locking, mutual exclusion etc)  

### Concurrency Vs Parallelism


A good concurrent program design is one where programming structure represents independently executing actions clearly. Parallelism is the ability of these well design concurrent program to run in parallel by providing a CPU core to run on. So it is important to remember that a well design concurrent program will also perform excellently when performing parallel capabilities.


By default the Go program runtime will run the main program assuming only one processor to be used. In such a case the go routines in the program will run currently but not in parallel. So by default one go routine runs at a time.


In order to give the go program to use more processors we can set the runtime property called GOMAXPROC() which gives the maximum processors to use.



### Go Routines and Channels

As mentioned earlier that go supports synchronization of the variables but does not promote it as it is inefficient way to do concurrency. However the go language uses the concept of channels or pipes that transport data between routines. At any time only one channel has access to the data. so no data race conditions happen at any time.

channel is defined as :  var channel_name chan <data_type>  Uninitialized channels are defaulted to value nil. Only data of the type that that is defined can be passed to the channel but we can send multiple data using interface{} as the type. It is useful to create a channel of channels.

The channels send and receive operations are atomic: they always complete without interruption.

```
var ch chan string
ch = make(chan string)
/**
  in short we can also write
  ch := make(chan string)
*/
// making a channel of channels
chOfChan := make (chan chan int)  

// make a channel of functions
chOfFunc := chan func()
```

#### Blocking of channels

By default the channels are synchronous and unbuffered: which means that the sender and receiver but must be read to send and receive. **Both sides are blocked util both are read to send and receive**

* A send operation is blocked until a receiver is available for the same channel. This is because a message sent on the channel is enough to fill it up and a the sender blocks until the receiver can consume the first message.
* A receiver blocks on the channel util the sender is ready and there is a value to consume from the channel.

Unbuffered channels make a perfect tool for synchronizing multiple go routines.

Pattern - using an unbuffered channel as a communication link between two go routines. The pattern can be seen in the routine_basics_test.go
