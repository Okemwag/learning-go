# Intermediate: Context

## Purpose

This lesson covers:

- what `context.Context` is
- values
- cancellation
- deadlines
- context cancellation in your own code

## What Is the Context?

`context.Context` is a standard Go interface used to carry:

- cancellation signals
- deadlines and timeouts
- small request-scoped values

It is one of the most important coordination tools in modern Go, especially in:

- HTTP handlers
- database calls
- background jobs
- concurrent workflows

The key idea is that context describes the lifecycle of a request or operation.

## Values

Context values let you attach small pieces of request-scoped metadata:

- request IDs
- trace IDs
- auth metadata

Use `context.WithValue` sparingly.

Good uses:

- metadata that must cross API boundaries
- values tied to the lifetime of the request

Bad uses:

- optional function arguments
- storing large data
- passing core business dependencies

If the value is really a parameter or dependency, pass it explicitly instead.

## Cancellation

`context.WithCancel` creates a child context that can be stopped explicitly.

Pattern:

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```

Then code can listen for:

```go
<-ctx.Done()
```

And check the reason with:

```go
ctx.Err()
```

This is the standard way to tell work to stop cooperatively.

## Contexts with Deadlines

Use:

- `context.WithDeadline`
- `context.WithTimeout`

These automatically cancel the context when time runs out.

This is the preferred pattern for:

- request timeouts
- bounded retries
- external calls with maximum wait times

Timeouts work best when lower-level code also respects the same context.

## Context Cancellation in Your Own Code

If your function:

- blocks
- loops
- waits on channels
- performs long-running work

then it should often accept a `context.Context` and check for cancellation.

A common shape is:

```go
func doWork(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case ...:
        ...
    }
}
```

This keeps your code responsive to:

- caller cancellation
- timeouts
- shutdown signals

## Practical Rules

- accept `context.Context` as the first parameter when appropriate
- do not store contexts inside structs unless there is a very specific reason
- do not pass `nil` contexts; use `context.Background()` or `context.TODO()`
- call the cancel function you create when you are done with it
- return `ctx.Err()` when cancellation or deadline ends the work

## Run It

```bash
go run ./intermediate/11_context
```

## Deep Note

Context is not a general-purpose bag of data. It is a lifecycle signal plus small request metadata. If you treat it that way, it becomes one of the cleanest ways to make cancellation, deadlines, and request propagation consistent across your program.
