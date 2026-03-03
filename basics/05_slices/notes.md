# Basics: Slices

## Purpose

This lesson covers:

- why slices are preferred over arrays
- `len`
- `append`
- capacity
- `make`
- emptying a slice
- declaring a slice
- slicing slices
- `copy`
- converting arrays to slices
- converting slices to arrays

## Why Slices Are Preferred

Slices are preferred because they are flexible:

- their length can grow and shrink
- they are cheap to pass around
- they work naturally with most Go APIs

Arrays are fixed-size values. Slices are lightweight descriptors with:

- a pointer to an underlying array
- a length
- a capacity

That makes slices the practical default for collections.

## Declaring a Slice

There are several common forms:

```go
var a []int
b := []int{1, 2, 3}
c := make([]int, 2, 5)
```

`var a []int` creates a nil slice.

## `len` and `cap`

- `len(slice)` is the number of elements currently visible
- `cap(slice)` is the amount of space available before a new allocation is required

Capacity matters because `append` may reuse the existing backing array or allocate a new one.

## `append`

`append` adds elements to a slice and returns the updated slice.

Always assign the result:

```go
s = append(s, 4)
```

## `make`

Use `make` for slices when you want to control initial size or capacity:

```go
s := make([]int, 2, 8)
```

That creates a slice of length 2 and capacity 8.

## Emptying a Slice

To reuse the backing array while clearing current contents:

```go
s = s[:0]
```

This is common in performance-sensitive loops.

## Slicing Slices

```go
window := s[1:4]
```

This creates a new slice view. It usually shares the same backing array.

That means modifying one view can affect another.

## `copy`

Use `copy(dst, src)` when you want a separate copy of values.

This is important when you need to avoid shared backing storage.

## Converting Arrays to Slices

An array can be sliced directly:

```go
arr := [3]int{1, 2, 3}
s := arr[:]
```

## Converting Slices to Arrays

You cannot directly cast an arbitrary slice to an array value in the general case.

The safe, common approach is:

- allocate an array
- copy slice data into it

That is what the example code demonstrates.

## Run It

```bash
go run ./basics/05_slices
```

## Deep Note

Most subtle collection bugs in Go come from shared slice backing arrays. If two slices overlap the same array, appending or mutating one can affect the other. Learn when you need `copy`, and the rest of slice behavior becomes much easier to reason about.
