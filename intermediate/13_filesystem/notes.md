# Intermediate: Filesystem APIs

## Purpose

This lesson covers the most common filesystem-related standard-library packages:

- `os`
- `path/filepath`
- `io/fs`

## `os`

Use `os` for direct operating-system interaction.

Common filesystem tasks:

- creating files and directories
- reading and writing files
- removing paths
- reading environment variables

High-value functions:

- `os.ReadFile`
- `os.WriteFile`
- `os.Open`
- `os.Create`
- `os.ReadDir`
- `os.MkdirAll`
- `os.RemoveAll`
- `os.MkdirTemp`

## `path/filepath`

Use `path/filepath` for local filesystem paths.

It handles OS-specific path separators and common path operations.

Useful functions:

- `filepath.Join`
- `filepath.Base`
- `filepath.Dir`
- `filepath.Ext`
- `filepath.Clean`
- `filepath.WalkDir`

Use `filepath` for local disk paths instead of `path`, which is more appropriate for slash-separated logical paths such as URLs.

## `io/fs`

`io/fs` provides filesystem abstractions and shared interfaces.

It becomes especially useful when code should work with:

- the real filesystem
- embedded files
- virtual filesystems

This keeps APIs flexible and testable.

## Practical Guidance

Prefer:

- `os.ReadFile` / `os.WriteFile` for simple whole-file operations
- `os.Open` plus readers/scanners for large streaming files
- `filepath.Join` instead of manual path concatenation
- temp directories in demos and tests to avoid polluting real paths

## Run It

```bash
go run ./intermediate/13_filesystem
```

## Deep Note

The filesystem APIs in Go are small but composable. Learn the distinction between direct OS calls (`os`), path manipulation (`filepath`), and abstract filesystem interfaces (`io/fs`), and a lot of file-handling code becomes much easier to structure cleanly.
