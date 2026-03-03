# Intermediate: `database/sql`

## Purpose

This lesson covers the standard-library database layer:

- `database/sql`
- `database/sql/driver`

## What `database/sql` Does

`database/sql` provides a common API for database access:

- opening connections
- pooling
- queries
- transactions
- scanning rows

It is not a database driver itself.

Instead:

- `database/sql` is the shared abstraction
- database-specific drivers plug into it

## Common Types

High-value types include:

- `sql.DB`
- `sql.Tx`
- `sql.Rows`
- `sql.Row`

Important idea:

- `sql.DB` is a pool handle, not one single connection

That means you usually keep one long-lived `*sql.DB` and reuse it.

## Common Operations

Typical patterns:

- `db.QueryContext(...)`
- `db.QueryRowContext(...)`
- `db.ExecContext(...)`
- `db.BeginTx(...)`

Prefer the `Context` variants in real applications so queries respect timeouts and cancellation.

## Scanning

Database rows are usually read into Go variables with `Scan(...)`.

Example:

```go
var name string
err := row.Scan(&name)
```

The scan destinations must be pointers.

## Driver Layer

`database/sql/driver` is the lower-level driver interface layer.

Most application code should not implement drivers, but seeing the boundary helps explain how `database/sql` works.

This lesson uses a tiny in-memory demo connector so the example is runnable without a real database server.

## Practical Guidance

Prefer:

- one shared `*sql.DB`
- context-aware query methods
- explicit error checking
- transactions for grouped writes

Be careful with:

- forgetting to close `Rows`
- opening and closing `sql.DB` repeatedly
- doing long-running work without context deadlines

## Real-World Note

In production, you pair `database/sql` with an actual driver such as:

- PostgreSQL drivers
- MySQL drivers
- SQLite drivers

This lesson stays driver-free so it remains runnable in a standalone learning repo.

## Run It

```bash
go run ./intermediate/15_database_sql
```

## Deep Note

`database/sql` is intentionally conservative: it gives you a stable abstraction for queries, transactions, and pooling, while leaving database-specific behavior to drivers. Learn the `sql.DB` lifecycle, context-aware queries, and proper scanning/closing habits first. Those are the habits that matter most in real systems.
