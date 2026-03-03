package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu      sync.Mutex
		counter int
		wg      sync.WaitGroup
	)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
