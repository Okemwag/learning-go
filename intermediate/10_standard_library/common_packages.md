# Commonly Used Standard Library Packages

This is a broader reference guide to the standard library packages you will use most often in everyday Go programs.

Use it alongside:

- [main.go](/home/okemwag/dev/GolangProjects/learning-go/intermediate/10_standard_library/main.go)
- [notes.md](/home/okemwag/dev/GolangProjects/learning-go/intermediate/10_standard_library/notes.md)

## Core Data and Formatting

### `fmt`

Use `fmt` for:

- printing
- formatted strings
- quick debugging output

Common functions:

- `fmt.Println(...)`
- `fmt.Printf(...)`
- `fmt.Sprintf(...)`
- `fmt.Errorf(...)`

This is one of the first packages every Go developer uses.

### `strings`

Use `strings` for immutable string processing.

Common operations:

- trimming
- splitting
- joining
- replacing
- prefix/suffix checks
- case conversion

Common functions:

- `strings.TrimSpace`
- `strings.Split`
- `strings.Join`
- `strings.Contains`
- `strings.HasPrefix`
- `strings.ToUpper`

### `bytes`

Use `bytes` when working with byte slices instead of strings.

It is useful for:

- binary data
- in-memory buffers
- efficient byte-oriented building

Common tools:

- `bytes.Buffer`
- `bytes.NewBuffer`
- `bytes.NewReader`
- `bytes.Equal`

### `strconv`

Use `strconv` for string conversion.

Common cases:

- string to int
- int to string
- parsing booleans
- formatting floats

Examples:

- `strconv.Atoi`
- `strconv.Itoa`
- `strconv.ParseBool`
- `strconv.FormatFloat`

## Input, Output, and Files

### `io`

`io` provides the core reader/writer interfaces used all over the standard library.

Most important interfaces:

- `io.Reader`
- `io.Writer`
- `io.Closer`
- `io.Seeker`

Common helper:

- `io.Copy`
- `io.ReadAll`

If you understand `io.Reader` and `io.Writer`, many Go APIs become much easier to understand.

### `bufio`

Use `bufio` to buffer I/O and work efficiently with streams.

Common tools:

- `bufio.NewReader`
- `bufio.NewWriter`
- `bufio.Scanner`

Use it for:

- reading line by line
- buffering file or network I/O
- token-style scanning

### `os`

Use `os` for operating-system interaction.

Common uses:

- opening files
- reading environment variables
- process info
- file and directory operations

Examples:

- `os.Open`
- `os.ReadFile`
- `os.WriteFile`
- `os.Getenv`
- `os.Create`

### `path/filepath`

Use `path/filepath` for filesystem paths.

It handles:

- OS-specific path separators
- joining paths
- walking directories
- extracting base names and extensions

Common functions:

- `filepath.Join`
- `filepath.Base`
- `filepath.Ext`
- `filepath.WalkDir`

Use this for local filesystem paths instead of `path`, which is usually for slash-separated paths such as URLs.

### `io/fs`

Use `io/fs` for filesystem abstractions.

It is useful for:

- embedded files
- virtual filesystems
- writing code against a filesystem interface

This package becomes especially relevant when using `embed`.

## Time, Context, and Concurrency

### `time`

Use `time` for:

- current time
- durations
- sleeping
- timers
- parsing and formatting timestamps

Common tools:

- `time.Now`
- `time.Since`
- `time.Sleep`
- `time.After`
- `time.NewTimer`
- `time.Parse`

### `context`

Use `context` to carry:

- cancellation
- deadlines
- request-scoped metadata

It is especially important in:

- HTTP handlers
- database work
- goroutine coordination

Common constructors:

- `context.Background`
- `context.WithCancel`
- `context.WithTimeout`
- `context.WithDeadline`
- `context.WithValue`

### `sync`

Use `sync` for shared-memory concurrency tools.

Common types:

- `sync.Mutex`
- `sync.RWMutex`
- `sync.WaitGroup`
- `sync.Once`
- `sync.Cond`
- `sync.Pool`

Use it when direct shared-state coordination is clearer than channels.

### `sync/atomic`

Use `sync/atomic` for low-level atomic operations on small shared values.

Common uses:

- counters
- flags
- tiny lock-free state transitions

Most application code should prefer:

- `sync.Mutex`
- channels

Atomics are a narrower tool than they first appear.

## Errors, Logging, and Diagnostics

### `errors`

Use `errors` for:

- creating sentinel errors
- joining errors
- checking wrapped errors

Common functions:

- `errors.New`
- `errors.Is`
- `errors.As`
- `errors.Join`

### `log`

Use `log` for simple application logging.

It is fine for:

- small tools
- quick scripts
- local development

For richer structured logging, prefer `log/slog`.

### `log/slog`

Use `log/slog` for structured logging.

It supports:

- leveled logs
- key/value fields
- pluggable handlers

This is the modern standard-library choice for application logging.

### `runtime` and `runtime/debug`

Use these packages when you need runtime information.

Common uses:

- goroutine counts
- stack traces
- build info
- GC diagnostics

These are more advanced, but they are important for debugging and observability.

## Encoding and Data Formats

### `encoding/json`

Use `encoding/json` for JSON encoding and decoding.

Common tools:

- `json.Marshal`
- `json.Unmarshal`
- `json.NewEncoder`
- `json.NewDecoder`

It is the standard default for API payloads and config-like JSON.

### `encoding/csv`

Use `encoding/csv` for CSV input and output.

Common tools:

- `csv.NewReader`
- `csv.NewWriter`

It is useful for imports, exports, and data pipelines.

### `encoding/base64`

Use `encoding/base64` when binary data must be represented as text.

Common cases:

- tokens
- binary payloads in JSON
- transport-safe encoded values

### `encoding/binary`

Use `encoding/binary` for binary formats and byte-order conversions.

It is common in:

- protocol handling
- file formats
- network data

This is usually safer than using `unsafe` for binary parsing.

## Networking and Web

### `net`

Use `net` for lower-level networking:

- TCP
- UDP
- listeners
- raw network connections

Examples:

- `net.Listen`
- `net.Dial`

### `net/http`

Use `net/http` for HTTP clients and servers.

It provides:

- `http.Client`
- handlers
- requests and responses
- middleware-friendly handler interfaces

This is one of the most important packages for backend Go.

### `net/url`

Use `net/url` for:

- parsing URLs
- query parameters
- URL encoding

Common tools:

- `url.Parse`
- `url.Values`

This is much safer than manual string concatenation for URLs.

## Data, Randomness, and Math

### `math`

Use `math` for floating-point math.

Examples:

- `math.Sqrt`
- `math.Round`
- `math.Abs`

### `math/rand`

Use `math/rand` for pseudo-random numbers.

Good for:

- simulations
- randomized tests
- non-security-sensitive random choices

Do not use it for security tokens.

### `crypto/rand`

Use `crypto/rand` for cryptographically secure random bytes.

Use it for:

- tokens
- secrets
- keys

If security matters, prefer this over `math/rand`.

## Collections and Helpers

### `sort`

Use `sort` for ordering slices.

It is useful for:

- stable output
- deterministic tests
- ranking and ordering

### `slices`

Use `slices` for modern generic slice helpers.

Common operations include:

- sorting
- cloning
- searching
- comparing

This package removes a lot of repetitive slice utility code.

### `maps`

Use `maps` for common generic map helpers.

It is useful for:

- cloning maps
- comparing maps
- extracting keys/values in helper workflows

This package complements the built-in map type.

## Processes, Configuration, and CLI Work

### `flag`

Use `flag` for simple command-line parsing.

It is good for:

- small CLI tools
- internal scripts
- demos

For larger CLIs, many teams use external libraries, but `flag` remains enough for many tasks.

### `os/exec`

Use `os/exec` to run external commands.

Common uses:

- integrating with other tools
- launching subprocesses
- capturing command output

Use it carefully because process management introduces failure modes and platform differences.

## Databases and Storage

### `database/sql`

Use `database/sql` as the standard database access layer.

It provides:

- connection pooling
- query execution
- transactions
- common DB abstractions

Actual database support comes from drivers, but `database/sql` is the standard surface.

## Text, Templates, and Parsing

### `regexp`

Use `regexp` for pattern matching.

It is powerful, but not always the clearest solution.

Prefer simpler string operations when they are enough.

### `text/template` and `html/template`

Use templates for text or HTML generation.

- `text/template` for generic text
- `html/template` for HTML with escaping

These are common for:

- emails
- reports
- server-rendered pages

## Package Selection Guidance

A practical default toolkit for many Go programs is:

- `fmt`
- `strings`
- `strconv`
- `errors`
- `io`
- `os`
- `path/filepath`
- `time`
- `context`
- `encoding/json`
- `net/http`
- `log/slog`
- `sync`

Those packages alone cover a large percentage of everyday Go work.

## How to Learn Them Effectively

A good way to learn the standard library is:

1. Learn the small interfaces first: `io.Reader`, `io.Writer`, `context.Context`.
2. Learn the common operational packages next: `time`, `errors`, `os`.
3. Learn the data and web packages after that: `encoding/json`, `net/http`.
4. Add concurrency and diagnostics tools as your programs grow.

## Deep Note

The most important part of the Go standard library is not the number of packages. It is how consistently the packages fit together. Once you understand the common interfaces and conventions, the rest of the library becomes much easier to navigate.
