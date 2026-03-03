# Intermediate: The Standard Library

## Purpose

This lesson covers:

- `io` and related interfaces
- `time`
- monotonic time
- timers and timeouts
- `encoding/json`
- struct tags as metadata
- unmarshaling and marshaling JSON
- JSON with readers and writers
- encoding and decoding JSON streams
- custom JSON parsing
- `net/http`
- the client
- the server
- `ResponseController`
- structured logging

For a broader package survey, also read:

- [common_packages.md](/home/okemwag/dev/GolangProjects/learning-go/intermediate/10_standard_library/common_packages.md)

## `io` and Friends

The `io` package provides a small set of interfaces that much of the standard library builds on.

The most important are:

- `io.Reader`
- `io.Writer`
- `io.Closer`

These interfaces are powerful because they decouple data sources from data consumers.

A function that accepts an `io.Reader` can work with:

- files
- network responses
- strings
- in-memory buffers

Likewise, a function that writes to an `io.Writer` can target:

- files
- HTTP responses
- buffers
- stdout

This is one of the standard library's strongest design patterns.

## Readers

A reader provides bytes on demand.

Common reader sources:

- `strings.NewReader(...)`
- `bytes.NewBuffer(...)`
- files from `os.Open(...)`
- HTTP response bodies

Readers make streaming and composable APIs possible.

## Writers

A writer accepts bytes.

Common writer targets:

- `bytes.Buffer`
- files
- `http.ResponseWriter`
- stdout/stderr

Many packages become easier once you think in terms of "read from here, write to there."

## `time`

The `time` package handles:

- timestamps
- durations
- sleeping
- timers
- parsing and formatting time values

Go code uses it constantly for deadlines, logging, retries, and scheduling.

## Monotonic Time

When you use `time.Now()` and later compute durations with:

- `time.Since(start)`
- `end.Sub(start)`

Go uses a monotonic clock component when available.

This matters because elapsed-time calculations remain stable even if wall-clock time changes.

That is why:

- use wall clock for display
- use monotonic-aware duration math for measuring elapsed time

## Timers and Timeouts

Useful tools:

- `time.NewTimer`
- `time.After`
- `context.WithTimeout`

Use:

- `time.After` for small local waits
- `context.WithTimeout` for request-scoped and API-scoped timeouts

In larger systems, `context` is usually the better timeout mechanism because it composes across calls.

## `encoding/json`

The `encoding/json` package converts between Go values and JSON.

The two core operations are:

- `json.Marshal`
- `json.Unmarshal`

It is the standard default for JSON handling in Go, especially for APIs and config data.

## Using Struct Tags to Add Metadata

Struct tags guide encoding behavior:

```go
type User struct {
    ID int `json:"id"`
}
```

This lets you:

- map Go field names to JSON field names
- control omission and formatting behavior

Struct tags are metadata consumed by reflection-based packages like `encoding/json`.

## Unmarshaling and Marshaling JSON

Marshaling:

```go
data, err := json.Marshal(value)
```

Unmarshaling:

```go
err := json.Unmarshal(data, &value)
```

Rules that matter:

- destination fields must be exported
- the destination should usually be a pointer
- custom types can implement JSON interfaces for custom behavior

## JSON, Readers, and Writers

`encoding/json` works naturally with streams:

- `json.NewDecoder(reader)`
- `json.NewEncoder(writer)`

This fits perfectly with the `io.Reader` / `io.Writer` ecosystem.

That means JSON can be parsed directly from:

- HTTP bodies
- files
- buffers
- network streams

without manually loading everything into a separate byte slice first.

## Encoding and Decoding JSON Streams

Stream-based JSON is useful when:

- data comes from an HTTP request body
- output goes to an HTTP response
- large payloads should not be fully buffered first
- multiple JSON values are encoded or decoded sequentially

This is often a cleaner and more memory-efficient approach than manual byte handling.

## Custom JSON Parsing

A type can implement:

- `json.Marshaler`
- `json.Unmarshaler`

This is useful when:

- a field needs special formatting
- external JSON format differs from your internal type
- you want stricter parsing rules

The lesson example uses a custom `Timestamp` type that controls how time is represented in JSON.

## `net/http`

`net/http` is the core HTTP package in Go.

It provides:

- clients
- servers
- handlers
- requests and responses

For many Go services, this package is the backbone of the application.

## The Client

A client sends requests and receives responses.

Basic pattern:

```go
client := http.Client{Timeout: ...}
resp, err := client.Get(url)
```

Important practices:

- set timeouts
- always close `resp.Body`
- check status codes explicitly

## The Server

A server usually starts from a handler:

```go
http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ...
})
```

The handler:

- reads the request
- writes the response
- sets status and headers

In real systems, handlers are often composed with middleware and dependency injection.

## `ResponseController`

`http.NewResponseController(w)` exposes lower-level response behavior for supported writers.

It can help with operations such as:

- flushing
- deadlines
- other response-level controls

This is useful in advanced handlers where the normal `ResponseWriter` surface is not enough.

Use it only when you need that lower-level control.

## Structured Logging

Structured logging means logs are emitted as:

- message + key/value fields

In modern Go, `log/slog` is the standard-library answer for structured logging.

Benefits:

- machine-readable fields
- better filtering and aggregation
- clearer operational context

Prefer:

- stable field names
- meaningful event messages
- small, useful contextual fields

Avoid dumping arbitrary large objects into logs unless necessary.

## Practical Guidance

These packages work especially well together:

- `io` provides the common data interfaces
- `encoding/json` consumes and produces readers/writers
- `net/http` exposes bodies and responses as readers/writers
- `time` controls deadlines and elapsed-time logic
- `slog` records structured operational events

Once you understand that ecosystem, a large part of practical Go backend work becomes much easier.

## Run It

```bash
go run ./intermediate/10_standard_library
```

## Deep Note

The standard library is one of Go's biggest strengths because the pieces fit together cleanly. `io.Reader` and `io.Writer` connect data flow, `time` manages deadlines safely, `encoding/json` builds on streaming interfaces, `net/http` uses the same abstractions for transport, and `slog` gives you operational visibility. Learning how these packages compose is more valuable than memorizing every individual function.
