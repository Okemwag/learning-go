# Intermediate: TCP Client

## Purpose

This example shows how network calls work at a lower level than `net/http`.

## What It Does

- opens a TCP connection with `net.Dial`
- sends a raw HTTP request
- reads the first response line from the server

## Run It

```bash
go run ./intermediate/03_tcp_client
```

## Deep Note

This is useful because it shows that HTTP is just a protocol running over a network connection. Understanding this makes higher-level HTTP tooling much easier to reason about.
