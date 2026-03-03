# Basics: Structs

## Purpose

This lesson covers:

- structs
- anonymous structs
- comparing structs
- converting structs

## Structs

A struct groups related fields into one value.

Example:

```go
type User struct {
    ID   int
    Name string
}
```

Structs are one of the most important tools in Go because they let you model data directly without needing heavyweight class systems.

## Anonymous Structs

An anonymous struct is a struct value declared without naming a reusable type.

Use anonymous structs when:

- the shape is small
- the data is local
- the type does not need reuse elsewhere

They are useful for:

- temporary grouped values
- tests
- JSON helper values

## Comparing Structs

Structs can be compared with `==` if all of their fields are comparable.

Comparable field examples:

- numbers
- strings
- booleans
- arrays of comparable elements

Non-comparable field examples:

- slices
- maps
- functions

If a struct contains a non-comparable field, the struct itself is not directly comparable.

## Converting Structs

Different struct types can sometimes be converted when their underlying field layout is compatible.

That is useful when:

- you have similar domain types
- you want explicit conversion between layers

Go keeps this explicit rather than automatic so that cross-type movement is intentional.

## Run It

```bash
go run ./basics/08_structs
```

## Deep Note

Structs are where Go starts feeling like a language for building real systems. Once you understand structs well, you can model requests, configuration, database rows, and business entities in a direct and readable way.
