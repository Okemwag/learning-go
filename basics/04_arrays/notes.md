# Basics: Arrays

## Purpose

This lesson focuses on arrays before moving on to slices.

## Arrays

An array in Go:

- has a fixed length
- stores elements of one type
- includes its length as part of its type

That last point matters:

- `[3]int` is a different type from `[4]int`

## Why Arrays Matter

Arrays are not used as often as slices in application code, but they are important because:

- slices are built on top of arrays
- arrays have value semantics
- arrays can be compared directly when elements are comparable

## Operations

- `len(array)` returns the number of elements
- indexing uses `array[i]`
- assignment copies the entire array

## Run It

```bash
go run ./basics/04_arrays
```

## Deep Note

Most Go developers use slices daily and arrays occasionally. Still, understanding arrays is essential because slice behavior only makes full sense when you know that a slice points into an underlying array.
