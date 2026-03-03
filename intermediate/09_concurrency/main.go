package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("when to use concurrency:")
	fmt.Println("- when work can happen independently")
	fmt.Println("- when latency can be hidden with overlapping I/O")
	fmt.Println("- not just because multiple cores exist")

	demoGoroutinesAndChannels()
	demoForRangeAndClose()
	demoBufferedVsUnbufferedAndBackpressure()
	demoSelectTimeoutAndDisableCase()
	demoLoopVariableCapture()
	demoWaitGroupAndOnce()
	demoMutexInsteadOfChannels()
	demoAtomicCounter()
	demoKeepAPIsConcurrencyFree()
}

func demoGoroutinesAndChannels() {
	// Goroutines are lightweight concurrent function executions.
	results := make(chan string)

	go func(out chan<- string) {
		out <- "hello from a goroutine"
	}(results)

	fmt.Println("goroutine result:", <-results)
}

func demoForRangeAndClose() {
	// A producer closes the channel when no more values will be sent.
	values := make(chan int)

	go func() {
		defer close(values)

		for i := 1; i <= 3; i++ {
			values <- i
		}
	}()

	// for-range over a channel keeps receiving until the channel is closed.
	for value := range values {
		fmt.Println("for-range channel value:", value)
	}

	// Reading from a closed channel gives the zero value plus ok=false.
	last, ok := <-values
	fmt.Println("closed channel read:", last, ok)
}

func demoBufferedVsUnbufferedAndBackpressure() {
	// Unbuffered channels synchronize sender and receiver immediately.
	unbuffered := make(chan string)
	go func() {
		unbuffered <- "unbuffered requires a receiver now"
	}()
	fmt.Println(<-unbuffered)

	// Buffered channels allow limited queued work before blocking.
	// This is useful for smoothing bursts, not for hiding unbounded load.
	jobs := make(chan int, 2)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for job := range jobs {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("processed job:", job)
		}
	}()

	// The capacity of 2 creates backpressure after two queued jobs.
	jobs <- 10
	jobs <- 20
	jobs <- 30
	close(jobs)
	<-done
}

func demoSelectTimeoutAndDisableCase() {
	fast := make(chan string, 1)
	slow := make(chan string, 1)

	go func() {
		time.Sleep(10 * time.Millisecond)
		fast <- "fast result"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		slow <- "slow result"
	}()

	// select waits on multiple channel operations.
	select {
	case msg := <-fast:
		fmt.Println("select chose:", msg)
	case msg := <-slow:
		fmt.Println("select chose:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("select timed out")
	}

	// Turning off a select case is usually done by setting the channel to nil.
	// nil channels are never ready, so that case is disabled.
	var optional <-chan string
	select {
	case msg := <-optional:
		fmt.Println("unexpected optional message:", msg)
	case <-time.After(5 * time.Millisecond):
		fmt.Println("optional case disabled with nil channel")
	}

	// Timeout code should usually use context in larger programs,
	// but time.After is fine for simple local waits.
}

func demoLoopVariableCapture() {
	var wg sync.WaitGroup

	// Always pass the loop variable as an argument.
	// Do not rely on the closure reading the changing loop variable later.
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			fmt.Println("safe loop variable:", value)
		}(i)
	}

	wg.Wait()
}

func demoWaitGroupAndOnce() {
	var (
		wg   sync.WaitGroup
		once sync.Once
	)

	start := func(workerID int) {
		defer wg.Done()

		// sync.Once ensures setup runs exactly once.
		once.Do(func() {
			fmt.Println("run once setup")
		})

		fmt.Println("worker started:", workerID)
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go start(i)
	}

	wg.Wait()
}

func demoMutexInsteadOfChannels() {
	// When shared state is simple, a mutex is often clearer than channels.
	var (
		mu      sync.Mutex
		counter int
		wg      sync.WaitGroup
	)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("mutex-protected counter:", counter)
}

func demoAtomicCounter() {
	// Atomics are for very small, specific shared-state operations.
	// Most application code does not need them directly.
	var counter atomic.Int64
	counter.Add(1)
	counter.Add(1)
	fmt.Println("atomic counter:", counter.Load())
}

func demoKeepAPIsConcurrencyFree() {
	// Keep public APIs synchronous and simple when possible.
	// Let the implementation decide whether concurrency is used internally.
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Millisecond)
	defer cancel()

	// This function looks like a normal synchronous call to the caller,
	// even though it uses a goroutine internally.
	values, err := collectValues(ctx, 3)
	if err != nil {
		fmt.Println("collectValues error:", err)
		return
	}

	fmt.Println("concurrency hidden behind normal API:", values)
}

func collectValues(ctx context.Context, max int) ([]int, error) {
	stream := make(chan int)

	go func() {
		defer close(stream)

		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		for i := 1; i <= max; i++ {
			select {
			case <-ctx.Done():
				// Always clean up your goroutines when cancellation happens.
				return
			case <-ticker.C:
				stream <- i
			}
		}
	}()

	var out []int

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case value, ok := <-stream:
			if !ok {
				return out, nil
			}
			out = append(out, value)
		}
	}
}
