# Intermediate: Types, Methods, and Interfaces

## Purpose

This lesson covers:

- types in Go
- methods
- pointer receivers and value receivers
- coding methods for nil instances
- methods as functions
- functions versus methods
- why type declarations are not inheritance
- types as executable documentation
- `iota`
- embedding for composition
- why embedding is not inheritance
- interfaces
- type-safe duck typing
- embedding and interfaces
- accept interfaces, return structs
- interfaces and `nil`
- interface comparability
- the empty interface
- type assertions and type switches
- using assertions and switches sparingly
- function types as a bridge to interfaces
- implicit interfaces and dependency injection
- Wire

## Types in Go

Types in Go are explicit and central to design.

They do several jobs:

- constrain valid operations
- communicate intent
- group behavior with methods
- make invalid states harder to represent

Go types are simple, but that simplicity is a strength. There is very little hidden behavior.

## Methods

Methods are functions attached to a type.

Example:

```go
func (c Counter) Snapshot() int
```

The receiver is the value before the method name.

Methods let you keep behavior close to the data it operates on.

## Pointer Receivers and Value Receivers

Use a value receiver when:

- the method does not need to mutate the receiver
- the type is small and cheap to copy
- value semantics are clearer

Use a pointer receiver when:

- the method mutates the receiver
- copying would be expensive
- the type should be treated as one shared object

Be consistent across a type when possible. Mixed receiver styles can be valid, but they should be deliberate.

## Code Your Methods for Nil Instances

Sometimes a nil receiver is a meaningful state, especially for pointer receiver methods.

This can be useful:

- when a nil value should behave as an empty object
- when you want defensive behavior instead of a panic

But do not do this automatically. Handle nil receivers only when it improves the API and keeps the behavior obvious.

## Methods Are Functions Too

In Go, methods can be treated as functions through:

- method expressions
- method values

Example:

```go
increment := (*Counter).Increment
```

That makes methods easier to pass around when needed.

## Functions Versus Methods

Use a method when:

- the behavior is tightly associated with the type
- the call should read like an operation on that value

Use a function when:

- the logic is broader utility behavior
- the operation is not owned by one type
- attaching the behavior to a type would be misleading

Methods communicate ownership of behavior. Functions communicate general utility.

## Type Declarations Are Not Inheritance

Creating a new type in Go does not mean "subclassing" another type.

For example:

```go
type UserID int
```

This creates a distinct type, not a child class of `int`.

There is no inheritance hierarchy here. It is a new named type with the same underlying representation.

## Types Are Executable Documentation

Named types add domain meaning:

- `UserID`
- `Status`
- `Message`

These names make code easier to read than raw primitives everywhere.

This is "executable documentation" because the meaning is not only written in comments. It is enforced by the type system.

## `iota` Is for Enumerations, Sometimes

`iota` is useful for related constants, especially enums:

```go
const (
    StatusDraft Status = iota
    StatusPublished
    StatusArchived
)
```

Use it when it improves clarity.

Do not use it when:

- explicit values are clearer
- stable external numeric meanings matter
- the constant group is not really an enum

`iota` is a tool, not a requirement.

## Use Embedding for Composition

Embedding lets one type include another type directly:

```go
type Document struct {
    AuditInfo
    Title string
}
```

This promotes the embedded type's fields and methods onto the outer type.

That is composition: building larger behavior from smaller pieces.

## Embedding Is Not Inheritance

Embedding is not subclassing.

It does not mean:

- "is-a" inheritance
- method overriding in the OOP sense
- Liskov-substitution-style class hierarchies

It means:

- "has-a" relationship
- field and method promotion
- reuse by composition

Think of embedding as structural convenience, not an inheritance tree.

## A Quick Lesson on Interfaces

Interfaces define behavior as method sets.

Example:

```go
type Speaker interface {
    Speak() string
}
```

Any type with that method satisfies the interface automatically.

## Interfaces Are Type-Safe Duck Typing

If a type has the required methods, it satisfies the interface. No explicit `implements` keyword is required.

That is duck typing with compile-time safety:

- flexible like duck typing
- checked by the compiler

This is one of Go's strongest design choices.

## Embedding and Interfaces

Interfaces can embed other interfaces:

```go
type ActiveSpeaker interface {
    Speaker
    Runner
}
```

This composes behavior requirements just like struct embedding composes data and methods.

## Accept Interfaces, Return Structs

This is one of the most useful design rules in Go.

Accept interfaces:

- makes inputs flexible
- allows easier testing
- reduces coupling

Return structs:

- preserves concrete capabilities
- avoids forcing callers into interface limitations too early
- keeps APIs more useful

Take behavior at the boundary, keep concrete power in the result.

## Interfaces and `nil`

An interface value is only `nil` when both are absent:

- no dynamic type
- no dynamic value

This is the classic trap:

```go
var r *Resource = nil
var s Speaker = r
```

Now `s != nil` because it still contains type information.

This confuses many Go developers early. Always be precise about whether you are checking:

- the interface value
- the concrete value inside it

## Interfaces Are Comparable

Interface values can be compared with `==` when their dynamic values are comparable.

That means:

- `any(10) == any(10)` is valid
- comparing interfaces holding slices or maps will panic at runtime because those concrete values are not comparable

So the rule is not "interfaces are always safe to compare." The actual rule is "interface comparison depends on the dynamic value."

## The Empty Interface Says Nothing

`interface{}` (or `any`) accepts any value.

That makes it flexible, but it communicates no behavior contract.

Use it when:

- you truly accept arbitrary values
- generic data plumbing is unavoidable
- reflection-like boundaries are involved

Avoid it when a real interface can express the needed behavior.

## Type Assertions and Type Switches

Type assertion:

```go
text, ok := value.(string)
```

Type switch:

```go
switch v := value.(type) {
case string:
    ...
}
```

These let you inspect dynamic interface values.

## Use Type Assertions and Type Switches Sparingly

They are sometimes necessary, but they often signal one of these problems:

- the interface is too vague
- `any` is being overused
- behavior should be expressed as methods instead

Prefer polymorphism over repeated type inspection when practical.

## Function Types Are a Bridge to Interfaces

A function type can implement an interface by adding a method:

```go
type FormatterFunc func(string) string
```

Then:

```go
func (f FormatterFunc) Format(s string) string
```

This pattern is very useful because it lets plain functions satisfy behavioral interfaces cleanly.

It shows how functions and interfaces can work together without heavy framework machinery.

## Implicit Interfaces Make Dependency Injection Easier

Go's implicit interfaces make dependency injection simple:

- define a small behavior interface
- write code against that interface
- pass in any concrete implementation

No container is required to benefit from DI.

This makes testing easier because test doubles can satisfy the same interface naturally.

## Wire

Wire is a compile-time dependency injection tool associated with the Go ecosystem.

The important concept is not the tool itself, but the wiring style:

- constructors declare dependencies explicitly
- dependencies are connected at build time
- no runtime service locator is required

In Go, manual wiring is often enough and is usually the best place to start.

Wire can help when construction graphs become large, but it should support clear constructors, not replace good design.

## Run It

```bash
go run ./intermediate/04_types_methods_interfaces
```

## Deep Note

Go's type system is powerful because it is restrained. Types communicate intent, methods attach behavior, embedding composes structure, and interfaces decouple behavior without inheritance. If you lean into those constraints instead of fighting them, your code stays easier to test, easier to wire together, and easier to understand.
