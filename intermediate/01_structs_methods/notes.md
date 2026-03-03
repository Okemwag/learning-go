# Intermediate: Structs and Methods

## Purpose

This example introduces custom types and methods.

## What to Notice

- `type User struct { ... }` defines structured data
- value receivers are useful when the method does not need to mutate data
- pointer receivers are useful when the method updates the struct

## Run It

```bash
go run ./intermediate/01_structs_methods
```

## Deep Note

Structs are the foundation of most Go programs. Once you understand structs and methods, you can model application state cleanly without heavy object-oriented patterns.
