# Intermediate: Modules, Packages, and Imports

## Purpose

This lesson covers:

- repository, modules, and packages
- `go.mod`
- the `go` directive
- the `require` directive
- building packages
- importing and exporting
- creating and accessing a package
- naming packages
- overriding a package's name
- Go doc comments
- `internal`
- avoiding circular dependencies
- organizing your module
- renaming and reorganizing APIs carefully
- avoiding `init` where possible
- importing third-party code
- versions and minimal version selection
- updating compatible and incompatible versions
- vendoring
- `pkg.go.dev`
- publishing and versioning a module
- overriding dependencies
- retracting versions
- workspaces
- proxy servers
- private repositories

## Repository, Modules and Packages

These are related but different:

- repository: the full source control project
- module: the versioned unit defined by a `go.mod`
- package: a directory of Go files compiled together

A single repository may contain:

- one module
- multiple modules
- many packages

In this repository, the root `go.mod` defines the main module.

## Using `go.mod`

`go.mod` is the module definition file.

It records:

- module path
- Go version intent
- required dependencies
- replacements
- retractions

Common commands:

```bash
go mod tidy
go get example.com/lib@latest
go list -m all
```

For most projects, keep one main `go.mod` at the repository root unless there is a real reason to split modules.

## Use the `go` Directive to Manage Go Build Versions

The `go` directive declares the language/toolchain version the module targets.

Example:

```go
go 1.25.1
```

This affects:

- language features available
- module graph behavior for some operations
- expectations for builds and tooling

It is not just documentation. It influences module behavior.

## The `require` Directive

`require` declares direct module dependencies and versions.

Example:

```go
require example.com/lib v1.2.3
```

Go tooling usually manages this for you. Prefer using:

- `go get`
- `go mod tidy`

instead of editing large dependency sets by hand.

## Building Packages

Build a package or command with:

```bash
go build ./intermediate/07_modules_packages
```

Or run it directly:

```bash
go run ./intermediate/07_modules_packages
```

Remember:

- packages compile by directory
- exported identifiers define the public surface

## Importing and Exporting

In Go:

- uppercase names are exported
- lowercase names are package-private

This applies to:

- functions
- types
- methods
- variables
- constants

The example package in this lesson demonstrates this in [greeting.go](/home/okemwag/dev/GolangProjects/learning-go/intermediate/07_modules_packages/greeting/greeting.go).

## Creating and Accessing a Package

Create a package by:

1. making a directory
2. adding Go files with the same package name
3. importing it by module path plus directory path

Example import in this lesson:

```go
import "github.com/Okemwwag/learning-go/intermediate/07_modules_packages/greeting"
```

## Naming Packages

Good package names are:

- short
- lowercase
- descriptive
- not repetitive with the import path

Prefer:

- `greeting`
- `store`
- `config`

Avoid:

- `utilities`
- `common`
- `helpers`

unless the package is genuinely that broad and unavoidable.

## Overriding a Package's Name

You can rename an import locally with an alias:

```go
import stdmath "math"
```

Reasons to do this:

- avoid name conflicts
- improve clarity at the call site
- shorten or disambiguate a package name

Use aliases sparingly. Most imports should keep their default name.

## Documenting Your Code with Go Doc Comments

Go documentation tools use leading comments.

Examples:

```go
// Package greeting demonstrates ...
package greeting
```

```go
// Build creates a formatted greeting.
func Build(name string) Message
```

Good doc comments:

- start with the identifier name
- describe behavior, not implementation trivia
- stay concise and factual

This is what tools like `go doc` and `pkg.go.dev` read.

## Using the `internal` Package

Any package inside an `internal` directory can only be imported by code within the parent tree.

This is a built-in visibility boundary for modules.

Use `internal` when:

- a package is shared inside your module
- you do not want external users depending on it
- you want freedom to refactor internals without breaking public consumers

This lesson includes an internal package at [secret.go](/home/okemwag/dev/GolangProjects/learning-go/intermediate/07_modules_packages/internal/secret/secret.go).

## Avoiding Circular Dependencies

Go forbids circular imports.

If package A imports B, and B imports A, the build fails.

Avoid this by:

- pushing shared types into a lower-level package
- defining interfaces at the consumer boundary
- keeping dependencies flowing in one direction

Circular dependencies are usually a design smell.

## Organizing Your Module

A practical module layout often separates:

- commands in `cmd/`
- internal app code in `internal/`
- reusable public packages only when truly needed

Keep the public API surface small. Most packages in a real app should not be designed for external reuse unless that is a real goal.

## Gracefully Renaming and Reorganizing Your API

When changing package layout or exported names:

1. keep old APIs temporarily when practical
2. add forwarding helpers or compatibility shims
3. update docs and examples
4. remove old paths in a planned version change

Avoid sudden breaking changes unless you are intentionally releasing a new incompatible major version.

## Avoiding the `init` Function if Possible

`init` runs automatically before `main`.

It can be useful, but it often hides setup and makes code harder to reason about.

Prefer:

- explicit constructors
- explicit registration calls
- clear setup in `main`

Use `init` only when the automatic behavior is truly justified.

## Working with Modules

Module workflow usually centers on:

- `go mod tidy`
- `go get`
- `go list -m`
- `go work` when multiple modules are involved

Let the toolchain manage dependency metadata whenever possible.

## Importing Third-Party Code

You typically add third-party code with:

```bash
go get example.com/lib@v1.2.3
```

Then import it in code.

Go will record the dependency in `go.mod` and `go.sum`.

## Working with Versions

Go modules use semantic versions.

Examples:

- `v1.2.3`
- `v2.0.0`

Understanding versioning matters because module paths and compatibility are tied together.

## Minimal Version Selection

Go uses Minimal Version Selection (MVS).

The key idea:

- the build chooses the minimum version of each module that satisfies the overall dependency graph requirements

This keeps version selection predictable and avoids some dependency-resolution complexity seen in other ecosystems.

## Updating to Compatible Versions

For compatible updates, you usually move within the same major version:

```bash
go get example.com/lib@latest
```

Or pin a specific compatible version.

Then run:

```bash
go mod tidy
```

## Updating to Incompatible Versions

In Go modules, incompatible major versions usually change the import path.

Example idea:

- `example.com/lib`
- `example.com/lib/v2`

That explicit path change makes breaking upgrades visible in code.

## Vendoring

Vendoring copies dependencies into a local `vendor/` directory.

Common command:

```bash
go mod vendor
```

Reasons to vendor:

- stricter reproducibility policies
- restricted network environments
- dependency auditing workflows

Many teams do not need vendoring day-to-day, but it remains useful in some environments.

## Using `pkg.go.dev`

`pkg.go.dev` is the primary public Go package documentation index.

Use it to:

- browse docs
- inspect exported APIs
- review examples
- see package metadata

It reads Go doc comments, so documentation quality directly affects how your package appears there.

## Publishing Your Module

To publish a module publicly:

1. push it to a supported VCS host
2. ensure the module path matches the repository path
3. tag versions using semantic version tags
4. keep docs and examples in good shape

Public visibility is easiest when the module path is stable and the repository is accessible.

## Versioning Your Module

Use semantic versioning:

- patch for fixes
- minor for backward-compatible additions
- major for breaking changes

In Go, major version changes from v2 onward usually require the major version suffix in the module path.

## Overriding Dependencies

Use `replace` in `go.mod` to override dependencies.

Common uses:

- local development against a local checkout
- temporary forks
- testing unpublished changes

Example:

```go
replace example.com/lib => ../local-lib
```

Use this carefully. Replacements are powerful, but they can create confusion if committed casually.

## Retracting a Version of Your Module

Use `retract` in `go.mod` when a published version should no longer be used.

This is useful for:

- broken releases
- accidental bad tags
- known severe issues

Retraction communicates that a version exists but should be avoided.

## Using Workspaces to Modify Modules Simultaneously

Go workspaces (`go.work`) let you work on multiple modules together without publishing temporary versions.

This is useful when:

- you are changing two local modules at once
- you want local edits to resolve together
- you are testing cross-module changes

Workspaces are for development coordination, not a replacement for clean module versioning.

## Module Proxy Servers

Go often downloads modules through proxy servers rather than hitting origin repositories directly.

This improves:

- caching
- reliability
- reproducibility

## Specifying a Proxy Server

Use the `GOPROXY` environment variable.

Example:

```bash
GOPROXY=https://proxy.golang.org,direct
```

This controls where module downloads are resolved.

## Using Private Repositories

Private modules usually require extra configuration such as:

- `GOPRIVATE`
- credentials for your VCS host
- sometimes custom proxy settings

The key idea is telling the Go toolchain which paths are private so it skips public proxy and checksum behavior where appropriate.

## Run It

```bash
go run ./intermediate/07_modules_packages
```

## Deep Note

Go modules and packages work well when the boundaries stay intentional: keep module structure simple, keep package APIs small, expose only what you want to support long-term, and let the Go toolchain manage dependency metadata instead of hand-editing it more than necessary.
