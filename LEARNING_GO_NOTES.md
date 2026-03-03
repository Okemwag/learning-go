# Learning Go Programming

This repository is now organized into a progressive learning path:

- [`basics/README.md`](/home/okemwag/dev/GolangProjects/learning-go/basics/README.md)
- [`intermediate/README.md`](/home/okemwag/dev/GolangProjects/learning-go/intermediate/README.md)
- [`advanced/README.md`](/home/okemwag/dev/GolangProjects/learning-go/advanced/README.md)

## How to Use This Repo

Each level contains:

- topic folders
- a `main.go` example inside each topic
- a matching `notes.md` file with deeper explanation

Run one topic directory at a time:

```bash
go run ./basics/01_hello
go run ./intermediate/01_structs_methods
go run ./advanced/01_goroutines_channels
```

## Learning Path

### Basics

Start here for:

- program structure
- predeclared types and zero values
- variables and constants
- arrays and slices
- strings, bytes, and runes
- maps
- structs
- blocks and control flow
- functions
- pointers and memory behavior

### Intermediate

Move here once the syntax feels natural:

- structs and methods
- interfaces
- types, methods, and composition
- generics
- errors and panic/recover
- modules, packages, and imports
- Go tooling
- concurrency
- standard library: io, time, json, http, slog
- context
- testing
- filesystem APIs
- templates
- database/sql
- explicit error handling
- networking with TCP

### Advanced

Finish with production-oriented patterns:

- goroutines and channels
- context cancellation and timeouts
- mutexes and shared state
- reflection, unsafe, and cgo guidance

## Suggested Workflow

For each topic:

1. Run the topic directory.
2. Read the matching `notes.md` file.
3. Modify the example and observe the result.
4. Rewrite the concept from memory in a new file.

## Core Commands

```bash
go run ./basics/03_variables_constants
go fmt ./...
go build ./...
go test ./...
```

## Build Note

The duplicate `main redeclared` errors are fixed by putting each example in its own package directory. That means `go test ./...` and `go build ./...` can now walk the repository cleanly.

Use `go run` on one topic directory at a time while learning.
