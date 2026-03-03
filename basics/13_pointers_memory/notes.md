# Basics: Pointers, Slices, Maps, and Memory

## Purpose

This lesson covers:

- pointers
- the difference between maps and slices
- slices as buffers
- reducing the garbage collector's workload
- tuning the garbage collector

## Pointers (Very Deep)

Pointers are values that store the address of another value.

Core syntax:

- `&x` gets the address of `x`
- `*p` reads or writes the value pointed to by `p`

Example:

```go
x := 10
p := &x
*p = 20
```

After that, `x` is now `20`.

## What Pointers Are For

Pointers are useful when:

- a function should modify the caller's value
- copying a value would be expensive
- a value needs to be shared intentionally
- a method should mutate a struct

Use pointers for clear reasons, not by default.

## What Pointers Are Not in Go

Go has pointers, but it deliberately avoids some low-level pointer features:

- no pointer arithmetic
- no manual memory free
- no direct stack vs heap control in user code

That makes pointers safer than in C, but they still require disciplined thinking.

## Pointer Semantics and Call by Value

Go is always call by value.

That includes pointers.

When you pass a pointer to a function:

- the pointer itself is copied
- both the original pointer and the copied pointer refer to the same underlying value

This explains two important facts:

1. Mutating `*p` inside the function changes the original pointed-to value.
2. Reassigning the local pointer variable does not change the caller's pointer variable.

That distinction is one of the most important pointer concepts in Go.

## Nil Pointers

A pointer's zero value is `nil`.

That means:

- it points to no value
- dereferencing it will panic

Always ensure a pointer is valid before dereferencing in code paths where `nil` is possible.

## Pointers to Structs

Pointers are especially common with structs:

```go
type User struct {
    Name string
}
```

Then:

```go
u := User{Name: "Ada"}
change(&u)
```

This allows mutation without copying the whole struct.

## The Difference Between Maps and Slices

Both maps and slices are passed by value, but their internal shapes differ.

### Slice

A slice is a small descriptor containing:

- a pointer to an underlying array
- a length
- a capacity

So when you pass a slice, Go copies that descriptor, not the entire underlying data.

Effects:

- changing an element through the slice usually affects the shared backing array
- changing the slice length in a called function does not update the caller's slice header unless you return the new slice
- appending may allocate a new array and break sharing

### Map

A map value refers to a runtime-managed hash table structure.

When you pass a map:

- the map value itself is copied
- both copies refer to the same underlying hash table state

Effects:

- inserting or updating entries inside a function is visible to the caller
- you usually do not need to return the map just to reflect content changes

### Practical Difference

This is the key practical distinction:

- slice element mutations are shared, but slice header changes are not automatically shared
- map content mutations are shared through the underlying map state

That is why slice APIs often return a slice, while map APIs often do not need to.

## Slices as Buffers

Slices are excellent reusable buffers.

A common pattern:

```go
buf := make([]byte, 0, 1024)
buf = append(buf, ...)
buf = buf[:0]
```

Why this matters:

- you reuse allocated memory
- you reduce repeated allocations
- you reduce pressure on the garbage collector

This is especially useful in:

- request processing
- serialization
- parsing
- logging pipelines

## Reducing the Garbage Collector's Workload

The garbage collector has more work when your program creates and discards many heap objects rapidly.

Common ways to reduce that workload:

### Reuse Buffers

Instead of allocating a new slice every time, reuse an existing one:

- reset with `s = s[:0]`
- keep a reasonable capacity

### Avoid Holding Large Backing Arrays by Accident

If a small slice still points into a huge backing array, the whole array may stay live.

Example problem:

- read a large buffer
- keep only a tiny subslice
- the tiny subslice keeps the large array from being reclaimed

A fix is to copy only the needed bytes into a smaller slice.

### Keep Object Lifetimes Short and Clear

Values that stop being referenced become collectible. Long-lived references increase retention.

Watch out for:

- global caches with no limits
- slices retaining huge backing arrays
- maps growing forever

### Reduce Unnecessary Pointer-Rich Structures

Pointer-heavy object graphs can increase GC scanning work.

Sometimes flatter data is cheaper:

- slices of values instead of slices of pointers
- compact structs
- fewer temporary heap objects

## Tuning the Garbage Collector

Most Go programs should start by writing simpler allocation-friendly code before touching GC tuning.

Tuning is a later step, not a first step.

### `GOGC`

The main GC tuning knob is `GOGC`.

It controls how much heap growth is allowed before the next GC cycle.

General idea:

- lower `GOGC` -> more frequent collection, lower memory use, potentially more CPU spent in GC
- higher `GOGC` -> less frequent collection, higher memory use, potentially less GC CPU overhead

Example environment variable:

```bash
GOGC=200 go run ./your-program
```

### When to Tune

Consider tuning only when you have evidence from:

- memory profiles
- allocation profiles
- latency measurements
- production metrics

If you tune blindly, you can easily make tradeoffs worse.

### Better First Steps Than Tuning

Before touching `GOGC`, usually do these first:

1. reduce allocations
2. reuse slices and buffers
3. avoid retaining large objects unnecessarily
4. measure with profiling tools

Those changes are often more valuable than runtime tuning alone.

## Run It

```bash
go run ./basics/13_pointers_memory
```

## Deep Note

Pointers, slices, maps, and the garbage collector are all connected by one core question: what data is still reachable, and what shape does that data have in memory? If you learn to reason about ownership, sharing, and retention, Go's performance characteristics become much more predictable.
