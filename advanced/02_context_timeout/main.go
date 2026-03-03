package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("operation completed")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := slowOperation(ctx); err != nil {
		fmt.Println("operation stopped:", err)
	}
}
