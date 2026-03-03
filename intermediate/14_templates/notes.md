# Intermediate: Templates

## Purpose

This lesson covers:

- `text/template`
- `html/template`

## `text/template`

Use `text/template` when generating plain text such as:

- emails
- reports
- config files
- generated source snippets

Templates support:

- field access
- conditionals
- loops
- functions

They are a clean way to separate formatting structure from the code that prepares data.

## `html/template`

Use `html/template` for HTML output.

Its key advantage is automatic escaping of untrusted content.

That makes it safer than `text/template` for web output because user-controlled strings are escaped by default.

## Practical Guidance

Use:

- `text/template` for non-HTML text
- `html/template` for HTML

Avoid:

- using `text/template` for HTML pages
- pushing too much business logic into templates

Templates should format prepared data, not become a second application layer.

## Run It

```bash
go run ./intermediate/14_templates
```

## Deep Note

The template packages are at their best when the Go code prepares clear data and the template only describes presentation. Keep logic shallow, keep data explicit, and let `html/template` handle escaping when output becomes web-facing.
