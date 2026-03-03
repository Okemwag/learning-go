# Advanced: Reflect, Unsafe, and Cgo

## Purpose

This lesson covers:

- reflection as runtime type inspection
- types, kinds, and values
- making new values with reflection
- checking whether an interface holds a nil value
- writing a data marshaler with reflection
- building functions with reflection
- why reflection cannot make methods
- when reflection is worth using
- why `unsafe` is unsafe
- `Sizeof` and `Offsetof`
- using `unsafe` with external binary data
- accessing unexported fields
- `unsafe` helper tools
- why cgo is for integration, not performance

## Reflection Lets You Work with Types at Runtime

The `reflect` package lets code inspect and manipulate values whose concrete types may not be known until runtime.

This is useful for:

- serialization frameworks
- ORMs
- generic adapters
- debugging and tooling

It is not the default choice for ordinary application logic.

Reflection is powerful, but it trades away:

- readability
- compile-time safety
- performance

## Types, Kinds, and Values

The key reflection concepts are:

- `reflect.Type`
- `reflect.Kind`
- `reflect.Value`

`Type` describes the named/static type information.
`Kind` describes the broad category such as:

- `struct`
- `slice`
- `map`
- `pointer`

`Value` represents the runtime value itself.

This distinction matters because:

- multiple named types can share the same kind
- operations often depend on kind, not just name

## Make New Values

`reflect.New(t)` creates a new zero value for type `t` and returns a reflective pointer.

This is useful when:

- building values dynamically
- decoding into runtime-selected types
- framework code needs to construct instances generically

Normal code should still prefer direct construction when the type is known.

## Use Reflection to Check if an Interface's Value Is Nil

One classic Go trap:

```go
var x any = (*User)(nil)
```

Now:

- `x != nil`

because the interface still contains type information.

Reflection can check the underlying value with:

- `reflect.ValueOf(x).IsNil()`

for kinds that support nil.

This is one of the few cases where reflection solves a real, practical problem cleanly.

## Use Reflection to Write a Data Marshaler

Reflection can iterate through struct fields and build a map or encoded form dynamically.

This is exactly how many serializers work.

It is useful when:

- the field list is not known statically
- generic metadata-driven output is needed
- building tooling, frameworks, or adapters

But the tradeoffs are real:

- slower than typed code
- easier to break
- more fragile under refactors

Use it when the generic behavior is genuinely worth that cost.

## Build Functions with Reflection, But Don't

`reflect.MakeFunc` can construct function values at runtime.

This is technically impressive, but usually a bad fit for normal code because:

- it is hard to read
- it is hard to debug
- it weakens compile-time clarity

If normal functions, closures, interfaces, or generics can express the same idea, use those instead.

## Reflection Can't Make Methods

Reflection can create values and functions, but it cannot change a type's method set at runtime.

Go types are fixed at compile time.

That means:

- you cannot add methods dynamically
- you cannot mutate a type definition at runtime

This is a deliberate language boundary that preserves predictability.

## Use Reflection Only If It's Worthwhile

Reflection is worthwhile when:

- the type truly is not known until runtime
- generic metadata-driven behavior is central to the feature
- the maintenance cost is justified

It is not worthwhile when:

- direct typed code would be simpler
- only one or two known types are involved
- the abstraction is just trying to be “generic” without real need

Most application code should use reflection sparingly.

## `unsafe` Is Unsafe

The `unsafe` package lets you bypass Go's type and memory safety rules.

That means you can:

- reinterpret memory
- compute offsets
- access otherwise restricted data

But you also give up important guarantees:

- portability
- safety
- future compatibility

Treat `unsafe` as a last-resort systems tool, not a convenience feature.

## Using `Sizeof` and `Offsetof`

Useful helpers:

- `unsafe.Sizeof(x)`
- `unsafe.Offsetof(field)`
- `unsafe.Alignof(x)`

These are most useful when:

- inspecting layout
- interoperating with binary formats
- validating assumptions in low-level code

They are informational tools, but they often accompany riskier `unsafe.Pointer` usage.

## Using `unsafe` to Convert External Binary Data

You can reinterpret raw bytes as a struct with `unsafe.Pointer`.

This can avoid copying, but it depends on assumptions about:

- endianness
- alignment
- struct layout
- architecture

That means the safer first choice is often:

- `encoding/binary`

Use `unsafe` only when you have measured the need and tightly controlled the input format.

## Accessing Unexported Fields

Reflection intentionally restricts access to unexported fields.

`unsafe` can bypass that restriction, for example by using:

- `Value.UnsafeAddr()`
- `reflect.NewAt(...)`

This is powerful, but it breaks encapsulation.

If you rely on it:

- code becomes fragile
- future implementation changes become dangerous
- maintenance becomes harder

This should be reserved for narrow tooling or compatibility layers, not routine application code.

## Using `unsafe` Tools

The main tools are:

- `unsafe.Pointer`
- `unsafe.Sizeof`
- `unsafe.Offsetof`
- `unsafe.Alignof`

These are low-level primitives. They should come with:

- strong comments
- narrow scope
- measured justification

If you use `unsafe`, leave a clear explanation of the invariant you rely on.

## Cgo Is for Integration, Not Performance

cgo exists to integrate Go with C libraries and native system APIs.

Good reasons to use cgo:

- calling a mature C library you must use
- binding to a platform-specific native API
- interoperating with existing non-Go components

Bad reasons to use cgo:

- chasing theoretical speed without profiling
- replacing straightforward Go code with C out of habit

cgo introduces real costs:

- more complex builds
- toolchain and platform dependencies
- boundary overhead between Go and C
- harder debugging and deployment

Use cgo when integration is necessary, not because it “sounds faster.”

## Run It

```bash
go run ./advanced/04_reflect_unsafe_cgo
```

## Deep Note

`reflect`, `unsafe`, and cgo all sit outside ordinary Go's comfort zone. They are useful precisely because they let you bypass normal boundaries, but that is also why they should be rare. Start with plain types, interfaces, generics, and the standard library. Reach for reflection only when runtime type work is truly required, `unsafe` only when the low-level tradeoff is justified, and cgo only when external integration demands it.
