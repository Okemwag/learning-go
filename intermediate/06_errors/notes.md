# Intermediate: Errors

## Purpose

This lesson covers:

- errors
- basic error handling
- string errors for simple examples
- sentinel errors
- errors as values
- wrapping errors
- wrapping multiple errors
- `Is` and `As`
- wrapping errors with `defer`
- `panic` and `recover`
- getting a stack trace from an error

## How to Handle Errors: The Basics

In Go, errors are handled explicitly.

The normal pattern is:

```go
value, err := doThing()
if err != nil {
    return err
}
```

This is the foundation of Go error handling:

- check errors directly
- return early on failure
- keep error flow visible

Go does not use exceptions for ordinary error handling.

## Use Strings for Simple Examples

For straightforward local failures, string-based errors are enough:

```go
return errors.New("name is required")
```

or:

```go
return fmt.Errorf("invalid input")
```

This is appropriate when:

- the error is simple
- no special programmatic handling is needed
- the message itself is enough

Do not over-engineer every error into a custom type.

## Sentinel Errors

Sentinel errors are shared package-level error values:

```go
var ErrNotFound = errors.New("not found")
```

They are useful when callers need to test for a specific condition:

```go
errors.Is(err, ErrNotFound)
```

Use sentinel errors for meaningful stable categories, not for every possible failure string.

## Errors as Values

This is one of Go's core ideas.

Errors are ordinary values:

- they can be returned
- wrapped
- compared
- stored
- placed inside structs

That means you can design them the same way you design other data:

- simple string values for simple cases
- structured custom types when callers need more detail

## Wrapping Errors

Wrapping preserves the original error while adding context:

```go
return fmt.Errorf("load user: %w", err)
```

This is important because it gives you both:

- a higher-level message
- access to the original cause

Use `%w` when you want the wrapped error to remain discoverable by `errors.Is` and `errors.As`.

## Wrapping Multiple Errors

Go can combine multiple failures with `errors.Join`:

```go
return errors.Join(err1, err2)
```

This is useful when:

- validating multiple fields
- collecting cleanup failures
- reporting independent errors together

Joined errors still work with `errors.Is` and `errors.As`.

## `Is` and `As`

Use `errors.Is` to check for a specific known error:

```go
errors.Is(err, ErrNotFound)
```

Use `errors.As` to extract a specific error type:

```go
var v ValidationError
if errors.As(err, &v) {
    ...
}
```

Rule of thumb:

- `Is` checks identity or semantic equivalence
- `As` extracts a typed error value

## Wrapping Errors with `defer`

Sometimes `defer` is the cleanest way to add final context:

```go
func f() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("f: %w", err)
        }
    }()
    ...
}
```

This is one of the few cases where a named return is justified:

- the deferred function needs access to the returned `err`

Use this pattern carefully. It is helpful, but less direct than ordinary explicit returns.

## `panic` and `recover`

`panic` is not normal error handling.

Use `panic` for:

- impossible states
- programmer bugs
- unrecoverable initialization failures

Do not use `panic` for:

- validation failures
- expected I/O errors
- normal business logic

`recover` can intercept a panic inside a deferred function:

```go
defer func() {
    if r := recover(); r != nil {
        ...
    }
}()
```

This is useful at process boundaries:

- server request boundaries
- worker loop boundaries
- framework safety layers

It should be rare in ordinary application logic.

## Getting a Stack Trace from an Error

A plain Go error does not automatically carry a stack trace.

If you need one, common approaches include:

- capturing `runtime/debug.Stack()` when creating a custom error
- using observability tooling that records stack context separately
- using panic recovery boundaries that log stacks

A simple custom pattern is:

```go
type StackError struct {
    Message string
    Stack   []byte
}
```

Then capture:

```go
debug.Stack()
```

This keeps the stack attached as structured data.

## Practical Guidance

Prefer this progression:

1. start with simple error returns
2. add wrapping for context
3. add sentinels only when callers need to branch on meaning
4. add custom types only when callers need structured inspection
5. use panic/recover only at true exceptional boundaries

This keeps error handling both explicit and maintainable.

## Run It

```bash
go run ./intermediate/06_errors
```

## Deep Note

The best Go error handling is direct, boring, and informative. Most functions should return ordinary errors with clear context. Reserve custom types, sentinels, joined errors, and panic/recover for the cases where they genuinely improve program behavior or diagnostics. If every error is "special," error handling quickly becomes harder to maintain than the business logic itself.
