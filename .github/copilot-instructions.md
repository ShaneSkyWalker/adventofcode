<!-- Repository-specific Copilot instructions for AI coding agents -->

# Copilot instructions — adventofcode

This repository contains a small, self-contained collection of Advent of Code solutions for 2025 written in Go. The guidance below focuses on patterns and workflows that make an AI agent immediately productive in this codebase.

- **Big picture**: each day is a standalone program under `2025/dayNN/` with a `main.go` that reads its input from `input.txt` (sometimes `input2.txt`) and implements `solvePart1` and `solvePart2`. The module name is `adventofcode25` (see `go.mod`). Example: [2025/day06/main.go](2025/day06/main.go#L1-L40).

- **How to run a single day**:

  - From repository root run:

    ```bash
    go run ./2025/day06
    ```

  - Or change directory and run:

    ```bash
    cd 2025/day06
    go run .
    ```

- **How to build or run all days**:

  - There is no umbrella `main` — run/build per-day. To build a single day binary:

    ```bash
    go build -o bin/day06 ./2025/day06
    ```

- **Tests**: Some days include ad-hoc test files (e.g. [2025/day02/test.go](2025/day02/test.go#L1-L40)). These are standalone `package main` helpers, not `*_test.go` unit tests. Use `go test ./...` only if you add real `_test.go` files.

- **Input handling pattern**:

  - Most `main.go` files define a `const inputFile = "input.txt"` and a local `readInput` helper that returns either `[]string` or multiple slices depending on the day's format. Inspect the day's `readInput` signature before refactoring. See [2025/day05/main.go](2025/day05/main.go#L1-L40) for an example that splits sections on blank lines.

- **Common conventions to follow**:

  - Keep each day's code self-contained in its folder; avoid introducing cross-day packages unless extracting genuinely reusable utilities (and then update `go.mod`).
  - Preserve existing `readInput` semantics for a day when modifying logic — callers expect the specific return shape.
  - Input file names: prefer `input.txt` for the main puzzle; `input2.txt` (when present) usually contains alternate/example input.

- **Formatting and style**:

  - Code is simple, imperative Go. Follow the existing style (no generics required). Keep functions small: `readInput`, `solvePart1`, `solvePart2`.

- **When adding fixes or features**:

  - Run the changed day locally with `go run` to validate output; do not modify `go.mod` unless adding dependencies.
  - If converting ad-hoc `test.go` helpers into proper tests, place them as `*_test.go` and rely on `go test`.

- **Files/directories to inspect for patterns**:

  - `go.mod` — module name and Go version.
  - `2025/dayNN/main.go` — per-day program layout; many days repeat the same input helpers.
  - `2025/dayNN/input.txt` and `input2.txt` — canonical inputs and examples.

- **Examples of quick edits an AI agent might be asked to perform**:

  - Refactor `readInput` into a shared helper only if at least two days share identical signatures and semantics.
  - Add a `Makefile` or top-level runner only if requested by the maintainer — otherwise keep per-day `go run` usage.

If anything here is unclear or you'd like different examples (running multiple days in parallel, CI steps, or converting helpers to shared packages), tell me which section to expand and I will update this file.
