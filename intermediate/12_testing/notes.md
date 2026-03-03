# Intermediate: Writing Tests

## Purpose

This lesson covers:

- understanding the basics of testing
- reporting test failures
- setting up and tearing down
- testing with environment variables
- storing sample test data
- caching test results
- testing your public API
- using `go-cmp` to compare results
- table tests
- running tests concurrently
- checking code coverage
- fuzzing
- benchmarks
- stubs
- `httptest`
- integration tests and build tags
- the data race detector

## Understanding the Basics of Testing

In Go, tests live in files ending with `_test.go`.

A test function looks like:

```go
func TestThing(t *testing.T) { ... }
```

Run tests with:

```bash
go test ./...
```

The `testing` package is built into the standard library, so writing tests requires no external framework.

## Reporting Test Failures

Common methods:

- `t.Error(...)`
- `t.Errorf(...)`
- `t.Fatal(...)`
- `t.Fatalf(...)`

Use:

- `t.Error` / `t.Errorf` when the test can continue checking more things
- `t.Fatal` / `t.Fatalf` when the test cannot continue safely

Failure messages should say:

- what happened
- what was expected

## Setting Up and Tearing Down

For per-test setup, prefer:

- local setup in the test body
- helper functions
- `t.Cleanup(...)`

For package-wide lifecycle setup, `TestMain(m *testing.M)` is available.

Use package-wide setup sparingly. Local setup is usually clearer.

## Testing with Environment Variables

Use `t.Setenv(...)` to set environment variables safely during a test.

It automatically restores the original value afterward.

This is better than calling `os.Setenv` manually because cleanup is automatic.

## Storing Sample Test Data

Go uses the `testdata/` convention for files that tests read.

The toolchain ignores `testdata/` when building normal packages, but tests can load files from it.

This is ideal for:

- fixture input files
- sample JSON
- golden files
- parser test cases

## Caching Test Results

`go test` caches successful test results for repeat runs when inputs have not changed.

This makes repeated local test runs much faster.

To force a fresh run:

```bash
go test -count=1 ./...
```

Caching is useful, but remember it exists when you are debugging a test that you expect to re-run.

## Testing Your Public API

Tests should primarily validate the public behavior of your package, not internal implementation details.

That makes refactoring easier because:

- implementation can change
- behavior stays stable

This lesson focuses tests on exported functions such as `Add`, `NormalizeName`, and `WelcomeUser`.

## Using `go-cmp` to Compare Test Results

For complex structs and nested values, `github.com/google/go-cmp/cmp` is often clearer than long manual comparisons.

Typical pattern:

```go
if diff := cmp.Diff(want, got); diff != "" {
    t.Fatalf("mismatch (-want +got):\n%s", diff)
}
```

This repository does not add that dependency in code, but it is worth knowing because it produces much clearer diffs than many hand-written assertions.

## Running Table Tests

Table tests are one of the most idiomatic Go testing patterns.

They:

- collect many cases into one structure
- reduce repetition
- make it easy to add cases

This lesson includes a table test for `NormalizeName`.

## Running Tests Concurrently

There are two related ideas:

- a test may start goroutines internally
- subtests can call `t.Parallel()`

Use parallel tests when:

- cases are independent
- shared state is controlled

Parallel tests can speed up suites, but only when the tests are isolated.

## Checking Your Code Coverage

Use:

```bash
go test -cover ./...
```

Or create a coverage profile:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

Coverage is useful, but it is not a quality guarantee. High coverage can still miss important behavior.

## Fuzzing

Go includes built-in fuzzing support.

A fuzz test looks like:

```go
func FuzzThing(f *testing.F) { ... }
```

Fuzzing is useful for:

- parsers
- string normalization
- boundary-heavy input handling

Run it with:

```bash
go test -fuzz=Fuzz ./intermediate/12_testing
```

## Using Benchmarks

Benchmarks measure focused code performance.

They look like:

```go
func BenchmarkThing(b *testing.B) { ... }
```

Run them with:

```bash
go test -bench=. ./intermediate/12_testing
```

Benchmarks help compare implementations, but they should measure realistic hot paths.

## Using Stubs in Go

A stub is a small fake implementation used to control a dependency in tests.

Go makes this simple when code depends on small interfaces.

This lesson uses a stub notifier to test behavior without a real external system.

This is one reason the advice "accept interfaces, return structs" is practical for tests.

## Using `httptest`

`net/http/httptest` is the standard way to test HTTP handlers and clients.

Useful tools:

- `httptest.NewRecorder()`
- `httptest.NewRequest(...)`
- `httptest.NewServer(...)`

Use `NewRecorder` for unit-style handler tests.
Use `NewServer` when you need a real HTTP server in an integration-style test.

## Using Integration Tests and Build Tags

Some tests are slower, broader, or require real infrastructure.

Keep those separate with build tags:

```go
//go:build integration
```

Then run them explicitly:

```bash
go test -tags=integration ./intermediate/12_testing
```

This keeps normal unit test runs fast while still preserving broader tests.

## Finding Concurrency Problems with the Data Race Detector

Use:

```bash
go test -race ./...
```

The race detector helps catch unsynchronized concurrent access to shared memory.

It is one of the highest-value test tools in Go.

Use it regularly for:

- concurrent code
- shared caches
- goroutine-heavy packages

It adds overhead, so it is slower, but it can reveal bugs that normal tests miss entirely.

## Practical Test Commands

Common commands:

```bash
go test ./...
go test -run TestNormalizeNameTable ./intermediate/12_testing
go test -cover ./...
go test -bench=. ./intermediate/12_testing
go test -race ./...
go test -fuzz=Fuzz ./intermediate/12_testing
```

## Run It

```bash
go test ./intermediate/12_testing
```

## Deep Note

Good Go tests are small, direct, and centered on behavior. Start with unit tests for the public API, add table tests for coverage of cases, use `httptest` for HTTP code, reserve integration tests for broader boundaries, and run the race detector on concurrent code. The best test suites are fast enough to run often and specific enough to fail for a useful reason.
