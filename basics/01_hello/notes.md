# Basics: Hello World

## Purpose

This is the minimum runnable Go program. It is intentionally small because you should understand the execution model before layering on types, collections, and structs.

## Core Ideas

- `package main` means this package builds as an executable program.
- `func main()` is the entry point. When the binary starts, Go calls this function.
- `fmt.Println(...)` is a standard library function that writes formatted text to standard output.

## Run It

```bash
go run ./basics/01_hello
```

## Deep Note

Go programs are organized into packages. A package can contain many files, but only a package named `main` can produce a directly runnable executable. Inside that package, the runtime looks for `main()`.

That simple structure is worth learning early because the rest of the language keeps the same philosophy:

- explicit structure
- clear entry points
- straightforward execution flow

Before moving on, be comfortable reading:

- package declarations
- import blocks
- function declarations
- standard library calls
