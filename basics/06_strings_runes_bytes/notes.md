# Basics: Strings, Runes, and Bytes

## Purpose

This lesson focuses on the relationship between:

- strings
- runes
- bytes

## Strings

A Go string is an immutable sequence of bytes.

That means:

- the data cannot be changed in place
- `len(s)` returns the number of bytes, not the number of human-readable characters

## Bytes

A byte is an alias for `uint8`.

Use bytes when dealing with:

- raw file data
- network payloads
- encodings
- binary protocols

Converting to `[]byte` gives you direct access to the underlying encoded data.

## Runes

A rune is an alias for `int32` and represents a Unicode code point.

Use runes when:

- counting characters in Unicode-aware logic
- transforming text by character
- iterating over multilingual text safely

## Why This Distinction Matters

Not every character is one byte.

For UTF-8 text:

- ASCII characters often use 1 byte
- many other characters use multiple bytes

That is why indexing and iterating can behave differently:

- `s[i]` gives a byte
- `for _, r := range s` gives runes

## Run It

```bash
go run ./basics/06_strings_runes_bytes
```

## Deep Note

Many beginners assume strings are arrays of characters. In Go, they are byte sequences. That choice makes I/O efficient and explicit, but it means you must choose between bytes and runes depending on what kind of text work you are actually doing.
