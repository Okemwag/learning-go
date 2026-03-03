# Advanced: Context and Timeouts

## Purpose

This example shows cancellation and deadlines with `context.Context`.

## What to Notice

- `context.WithTimeout` sets a deadline
- `defer cancel()` releases resources promptly
- long-running work should listen for `ctx.Done()`

## Run It

```bash
go run ./advanced/02_context_timeout
```

## Deep Note

In real services, `context.Context` is how request cancellations, timeouts, and shutdown signals flow through the system. It is a core skill for backend Go.
