# Intermediate: Interfaces and Errors

## Purpose

This example combines two core Go ideas:

- interfaces define behavior
- errors are ordinary values

## Interfaces

`Speaker` does not care about concrete type details. Any type with `Speak() string` satisfies the interface automatically.

This is one of Go's strongest design choices:

- no explicit `implements` keyword
- small interfaces are preferred

## Errors

`normalizeName` returns `(string, error)`.

That pattern is standard in Go because it keeps failure states visible and easy to handle.

## Run It

```bash
go run ./intermediate/02_interfaces_errors
```
