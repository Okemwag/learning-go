package main

import (
	"context"
	"fmt"
	"time"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func main() {
	// What is the context?
	// context.Context carries request-scoped cancellation, deadlines,
	// and small pieces of request-scoped metadata.
	base := context.Background()

	// Values:
	// context values are for request-scoped metadata that crosses API boundaries.
	ctxWithValue := context.WithValue(base, requestIDKey, "req-123")
	fmt.Println("context value:", requestIDFrom(ctxWithValue))

	// Cancellation:
	// WithCancel creates a derived context that can be cancelled explicitly.
	cancelCtx, cancel := context.WithCancel(ctxWithValue)
	done := make(chan struct{})

	go func() {
		defer close(done)
		waitForCancellation(cancelCtx)
	}()

	time.Sleep(5 * time.Millisecond)
	cancel()
	<-done

	// Contexts with deadlines:
	// WithTimeout is a common shorthand for deadline-based cancellation.
	deadlineCtx, deadlineCancel := context.WithTimeout(base, 15*time.Millisecond)
	defer deadlineCancel()

	if err := doTimedWork(deadlineCtx); err != nil {
		fmt.Println("deadline result:", err)
	}

	// Context cancellation in your own code:
	// your function should select on ctx.Done() when it can block or run for a while.
	customCtx, customCancel := context.WithTimeout(ctxWithValue, 25*time.Millisecond)
	defer customCancel()

	values, err := generateSequence(customCtx, 10)
	if err != nil {
		fmt.Println("generateSequence error:", err)
		return
	}
	fmt.Println("generated before cancellation:", values)
}

func requestIDFrom(ctx context.Context) string {
	value, _ := ctx.Value(requestIDKey).(string)
	return value
}

func waitForCancellation(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("cancellation received:", ctx.Err())
}

func doTimedWork(ctx context.Context) error {
	select {
	case <-time.After(30 * time.Millisecond):
		fmt.Println("timed work finished")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func generateSequence(ctx context.Context, max int) ([]int, error) {
	var out []int

	for i := 1; i <= max; i++ {
		select {
		case <-ctx.Done():
			// Your own code should stop promptly and return the context error.
			return out, ctx.Err()
		case <-time.After(10 * time.Millisecond):
			out = append(out, i)
		}
	}

	return out, nil
}
