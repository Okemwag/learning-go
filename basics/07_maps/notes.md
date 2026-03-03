# Basics: Maps

## Purpose

This lesson covers:

- maps
- reading and writing a map
- the comma-ok idiom
- deleting from maps
- emptying a map
- comparing maps
- using maps as sets

## What a Map Is

A map stores key-value pairs.

Example:

```go
map[string]int
```

This means:

- keys are `string`
- values are `int`

## Reading and Writing

Write:

```go
scores["alice"] = 90
```

Read:

```go
value := scores["alice"]
```

If the key is missing, Go returns the value type's zero value.

## The Comma-Ok Idiom

Because missing keys return zero values, you often need to know whether the key actually exists.

Use:

```go
value, ok := scores["alice"]
```

- `value` is the stored value or zero value
- `ok` is `true` if the key exists

## Deleting from Maps

```go
delete(scores, "alice")
```

Deleting a missing key is safe and does nothing.

## Emptying a Map

A common pattern is:

```go
for key := range scores {
    delete(scores, key)
}
```

You can also replace the map with a new one if that fits the code better.

## Comparing Maps

Maps cannot be compared with `==` except against `nil`.

That means this is invalid:

```go
// scores1 == scores2
```

If you need equality, compare contents manually or use a helper such as `maps.Equal` in newer Go versions.

## Using Maps as Sets

Go has no built-in set type, so maps are commonly used as sets.

Typical patterns:

- `map[string]bool`
- `map[string]struct{}`

`map[T]struct{}` is memory-efficient because the value carries no data.

## Run It

```bash
go run ./basics/07_maps
```

## Deep Note

Maps are reference-like runtime structures. They are fast and convenient, but iteration order is not stable. Never write logic that depends on map iteration order unless you sort keys first.
