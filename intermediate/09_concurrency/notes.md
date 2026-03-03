# Intermediate: Concurrency

## Purpose

This lesson covers:

- when to use concurrency
- goroutines
- channels
- reading, writing, and buffering
- `for-range` with channels
- closing a channel
- how channels behave
- `select`
- concurrency practices and patterns
- keeping APIs concurrency-free
- goroutines, `for` loops, and varying variables
- cleaning up goroutines
- buffered vs unbuffered channels
- backpressure
- turning off a `select` case
- timeout code
- `WaitGroup`
- `sync.Once`
- combining concurrency tools
- when to use mutexes instead of channels
- atomics

## When to Use Concurrency

Use concurrency when:

- work can proceed independently
- latency can be reduced by overlapping I/O
- background processing improves responsiveness
- coordination between tasks is genuinely needed

Do not use concurrency just because:

- the problem looks "advanced"
- you want to appear faster without measuring
- a simple synchronous design would be clearer

Concurrency is a tool for coordination and throughput, not a default programming style.

## Goroutines

A goroutine is a lightweight concurrent function execution started with:

```go
go doWork()
```

Goroutines are cheap, but not free. Each goroutine should have:

- a clear purpose
- a clear owner
- a clear shutdown path

## Channels

Channels are typed conduits used to send values between goroutines.

Example:

```go
ch := make(chan int)
```

Common operations:

- `ch <- value` sends
- `value := <-ch` receives
- `close(ch)` closes the channel

Channels help coordinate work and communicate ownership of data.

## Reading, Writing, and Buffering

Channel operations are synchronous unless buffering allows temporary queueing.

- send: `ch <- value`
- receive: `value := <-ch`
- buffered channel: `make(chan T, n)`

Buffered channels let a limited number of values queue without an immediate receiver.

Unbuffered channels require sender and receiver to meet at the same moment.

## Using `for-range` and Channels

You can iterate over a channel with:

```go
for value := range ch {
    ...
}
```

This continues until the channel is closed.

This is one of the cleanest ways to consume a stream of work from a producer.

## Closing a Channel

Closing a channel signals:

- no more values will be sent

It does not mean:

- "drop all queued values"
- "the receiver should stop immediately without draining"

Important rule:

- the sending side should usually be the side that closes the channel

Closing a channel from the wrong side is a common source of panics and races.

## Understanding How Channels Behave

Important channel facts:

- sending to a nil channel blocks forever
- receiving from a nil channel blocks forever
- sending to a closed channel panics
- receiving from a closed channel returns the zero value immediately after buffered values are drained
- receiving from a closed channel can report `ok == false`

These behaviors matter because many `select` patterns rely on them deliberately.

## Channel Direction

Channel parameters can express intent:

- `chan<- T` means send-only
- `<-chan T` means receive-only

This is useful because it narrows what a function is allowed to do and makes concurrency code easier to reason about.

## `select`

`select` waits on multiple channel operations and runs one ready case.

Example:

```go
select {
case msg := <-fast:
    ...
case <-time.After(time.Second):
    ...
}
```

Use `select` when:

- you need timeouts
- multiple channels may produce work
- cancellation must interrupt waiting
- you want to disable or enable cases dynamically

`select` is a core building block for robust concurrent systems.

## Turn Off a Case in a `select`

A common pattern is setting a channel variable to `nil`.

Because nil channels are never ready:

- the case becomes disabled

This is a clean way to remove a branch from an active `select` loop without duplicating the entire control structure.

## Timeout Code

Simple timeout:

```go
select {
case msg := <-ch:
    ...
case <-time.After(time.Second):
    ...
}
```

For larger systems, prefer `context.Context` so timeouts and cancellation integrate across function boundaries.

## Cancellation with `context`

The standard way to stop goroutines cooperatively is `context.Context`.

Typical pattern:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

Then the goroutine listens for:

```go
<-ctx.Done()
```

This avoids ad-hoc shutdown flags and keeps cancellation consistent across the program.

## Concurrency Practices and Patterns

Prefer:

- one clear owner per goroutine lifecycle
- explicit shutdown paths
- small, narrow channels
- context-aware cancellation for long-running work
- measured concurrency instead of guesswork

Avoid:

- "fire and forget" goroutines with no cleanup path
- hidden blocking points
- using concurrency in APIs when callers do not need to know about it
- mixing too many concurrency mechanisms without a clear reason

## Keeping Your APIs Concurrency-Free

A strong design pattern is:

- keep the public API synchronous and simple
- use goroutines internally if the implementation benefits from them

This keeps callers from having to reason about:

- channel lifecycles
- goroutine cleanup
- partial completion states

Concurrency should often be an implementation detail, not part of the public contract.

## Goroutines, `for` Loops, and Varying Variables

Loop variable capture is a classic mistake.

Bad pattern:

- a closure references a loop variable that keeps changing

Safer pattern:

- pass the current loop value as a function argument

This makes each goroutine receive the exact intended value.

## Always Clean Up Your Goroutines

Every goroutine should eventually stop unless it is intentionally process-long.

Use:

- `context`
- channel closure
- done channels
- clearly bounded loops

Goroutine leaks are a real production issue. They often look harmless in small examples and become expensive later.

## Know When to Use Buffered and Unbuffered Channels

Use unbuffered channels when:

- you want direct handoff
- synchronization between sender and receiver matters

Use buffered channels when:

- a small queue is useful
- you want to absorb short bursts
- the sender should proceed briefly without an immediate receiver

Do not treat buffering as a fix for poorly designed concurrency. It changes timing, not fundamentals.

## Implement Backpressure

Backpressure means slowing producers when consumers cannot keep up.

Buffered channels are a simple backpressure tool:

- once the buffer fills, sends block

This prevents unbounded growth and forces load to be absorbed where capacity actually exists.

Backpressure is usually better than silently accumulating unlimited work in memory.

## Use WaitGroups

`sync.WaitGroup` is for waiting until a known set of goroutines finishes.

Use it when:

- you start several workers
- the current function must wait for all of them to complete

It is one of the simplest and most useful concurrency coordination tools.

## Run Code Exactly Once

`sync.Once` ensures initialization or setup runs only once, even under concurrent access.

Use it for:

- one-time initialization
- lazy setup
- idempotent shared startup logic

It is clearer and safer than manually protecting a "did this happen already?" flag.

## Put Your Concurrent Tools Together

Real Go concurrency often combines:

- goroutines for work
- channels for communication
- `select` for multiplexing
- `context` for cancellation
- `WaitGroup` for joining
- `sync.Once` for one-time setup

The goal is not to use all of them at once. The goal is to combine them only when each solves a specific coordination problem.

## When to Use Mutexes Instead of Channels

Use a mutex when:

- you have simple shared state
- direct locking is clearer than message passing
- the state is small and local

Use channels when:

- values are being handed off between goroutines
- coordination is more about sequencing than shared memory
- ownership transfer is important

A mutex is often the better choice for a counter, cache map, or small shared struct.

## Atomics (Why You Probably Don't Need These)

Atomics are for very small, low-level shared-state operations such as:

- counters
- flags
- tiny performance-sensitive primitives

Most application code should prefer:

- channels
- mutexes

Why you probably do not need atomics often:

- they are easy to misuse
- they solve narrower problems than they appear to
- they make invariants harder to reason about when multiple fields are involved

If the state is more complex than a single value, a mutex is usually clearer.

## Practical Guidance

Prefer:

- channels for coordination and ownership transfer
- `context` for cancellation and deadlines
- clear goroutine lifecycles
- mutexes for simple shared state

Be careful with:

- goroutine leaks
- channels that are never closed when callers expect closure
- blocking sends with no receiver
- complex concurrency without a clear shutdown path
- unnecessary atomics

## Run It

```bash
go run ./intermediate/09_concurrency
```

## Deep Note

Good Go concurrency is not about spawning more goroutines. It is about controlling lifecycle and pressure: who owns the work, where values wait, what blocks, when everything stops, and how errors or cancellation propagate. If those answers are unclear, the concurrency code is not ready yet.
