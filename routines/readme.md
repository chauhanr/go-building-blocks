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

* **Pattern 1: Unbuffered Channel Go Routine Pattern** - using an unbuffered channel as a communication link between two go routines the sender immediately blocks as soon as the first message is sent on the channel. The pattern can be seen in the routine_chan_usage_patterns_test.go under the TestUnbufferedChannelPattern test case.

* **Pattern 2: Buffered Synchronous Channel Pattern** - In the buffered asynchronous pattern of communication channel has a buffer size and will accept value from the sender in a non blocking mode till the buffer is full. As soon as the buffer fills up the sender is blocked. This pattern is often used to make the first pattern more scalable.This pattern can be seen in the TestAsyncChannelsWithBufferPattern in the routine_chan_usage_patterns_test.go test file.

* **Pattern 3: [Semaphore] Using channels as output pipe** - This pattern is used to put a algorithm that is processing intensive into its own routine and passing a channel to it. the output of the function is returned back to the calling method using a channel. The channel is then read from the calling method which blocks till the routine returns. The pattern can be seen in the TestChannelAsOutputRoutine test case in the routine_chan_usage_patterns_test.go file. This Pattern is a semaphore pattern.  

```
// example of a buffered semaphore channel.
type Empty interface{}
var empty Empty
...
data := make([]float64, N)
res := make([]float64, N)
sem := make(chan Empty, N)
...

for i, ix := range data{
    go func (i int, xi float64){
        res[i] = doSomething(i,xi)
        sem <- empty
    }(i,ix)
}

// wait for the semaphores
for i :=0; i<N, i++{
    <-sem
}

```
* **Pattern 4: Running parallel for loop** - The example snippet above shows how a each iteration in the for loop can be converted to a go routine. This allows for each iteration to be spawned into a separate routine each of which could potentially get a processor to work with. TestChannelAsOutputRoutine test case in the routine_chan_usage_patterns_test.go file also has this pattern where we call a go routine for each for loop run.


* **Pattern 5: Channel Factory Pattern** - In this pattern the channel is not created by the client but is generated by the producer and then piped into the consumer so that client does not need to manage the channel. It is the job of the producer to make the channel and the consumer to consume from it. This can also be seen as a Producer Consumer pattern.

```

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

// calling method will call the two above as

suck(pump(9))

```

* **Pattern 6: Channel Iterator Pattern** - If we have a container that has some items in it that need to be processed by a consumer but we want this processing to happen in a separate thread we can do so by following the iterator pattern here:

```
  func (c *container) Iterator() <- chan {
      ch := make(chan item)
      go func(){
          for i :=0; i< c.Len(); i++{
            ch <- c.items[i]
          }
      }()   
      return ch
  }


  // the calling method can now iterate through the channel in a separate go routine.

    for x := range container.Iterator() {
      ...
    }

```

* **Pattern 7: Pipes and Filter Pattern** - pipes and filter pattern is also a very popular and heavily popular pattern. Here two channels work in tandem one is the in channel and the other is out channel. The generateNumbers(), filter() and sieve() functions in the routines_basics.go file are examples of using pipes and filters with channels.

The pipes and filters pattern has to be used with two patterns of closing the channel when it is no longer being used. this can be done by using the close() function in go as well as the boolean return value that a channel returns.

```
    value, ok := <-chan
    if !ok {
      break // break from the looping function that is reading from the channel.
    }else{
      process(value)
    }
```

if the close and ok operation then we will get deadlock errors where the channels are all waiting for each other.

* **Pattern 8: Switch statement with Channels** - there is a special switch statement syntax called select that works with the channels and go routines.

```
  for {
      select{
        case u, oku := <- chana:
              if !oku {
                break
              }
              // process the channel element u
              process(u)

        case v, okv := <- chanb:
                  if !okv {
                    break
                  }
                  // process the channel element v
                  process(v)
        default: // this is optional if present will cause the select to the non blocking.
      }
  }
```

Select stops once a break or return is encountered in one of its cases. Here is the behavior of select statement in various scenarios.
  * if all channels are blocked, it waits until one can processed
  * if multiple can proceed then select statement will choose one at random.
  * when none of the cases or channels can proceed and default clause is present, then default is executed: the default is always runnable. A default clause ensures that the select statement is non blocking.

The select pattern above is very close to how server backend process requests and here is a code that can mimic server process.
```
   func backend(){
      for {
          select {
            case cmd := <-ch1
                   // handle command on channel 1
            case cmd := <-ch2
                  // handle command on channel 2
            case cmd := <-chStop
                  // handle server stop command.
          }
      }
   }
```

* **Pattern 9:Timeout patterns** There are a few patterns that can help with timeouts when dealing with go routines.
    * Simple timeout pattern - have a lambda go routine that has a sleep for the time out period and then sending signal on channel.

```
  timeoutChannel := make(chan bool, 1)
  go func(){
    time.Sleep(1e9)
    timeout <- true
  }()

  // then the timeout can be handled by the select statement.
  select{
    case <-ch:
        // read from the ch has occurred
    case <-timeout:
       break
  }
```

    * Abandon the sync calls that run too long.
```
  ch := make(chan error, 1)
  go func(){ ch<- client.Call(args, &reply)}()

  select{
    case resp := <-ch:
        // use the resp and reply.
    case <-time.After(timeoutN):
        break
  }
```
      * Choose the function that returns the value the fastest. - in the function given below we have connections to multiple databases and we need to run the query on each one of them and return the query result that finishes first.

```
    func Query(conns []Conn, query string) Result{
      ch := make(chan Result, 1)
      for _, conn := range conns{
        go func(c Conn){
          select{
              case ch <- c.DoQuery(query)
              default:
          }
        }(conn)
      }
      return <-ch
    }
```

**Recovering from routines which error out** - if you want to only recover from a panic or error in a routine and not kill other routines we need to use the recover() method in golang.

```
  func server(workChan <-chan *Work){
      for work := range workChan{
        go workSafely(work)
      }    
  }

  func workSafely(work *Work){
    defer func(){
        if err := recover(); err != nil{
          fmt.Printf("Error %s while executing worker %v:\n", err, work)
        }
    }
    do(work)
  }
```
