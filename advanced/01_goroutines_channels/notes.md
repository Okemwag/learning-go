# Advanced: Goroutines and Channels

## Purpose

This example introduces concurrency using goroutines and channels.

## What to Notice

- `go worker(...)` starts work concurrently
- receive-only and send-only channel types make intent explicit
- buffered channels can hold values without blocking immediately

## Run It

```bash
go run ./advanced/01_goroutines_channels
```

## Deep Note

Concurrency is where Go becomes especially valuable for servers and systems work. Learn how channel ownership and channel direction reduce synchronization mistakes.
