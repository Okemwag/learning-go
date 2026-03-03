# Advanced: Mutex and Shared State

## Purpose

This example shows how to protect shared memory with `sync.Mutex`.

## What to Notice

- multiple goroutines update the same variable
- a mutex guards the critical section
- `sync.WaitGroup` waits for all workers to finish

## Run It

```bash
go run ./advanced/03_mutex_shared_state
```

## Deep Note

Go gives you both channels and locks. Use channels for coordination and ownership transfer. Use a mutex when shared state is simple and direct locking is clearer.
