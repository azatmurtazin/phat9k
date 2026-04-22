# Project Context

**Phat9k** - PHP to Go transpiler, parser, analyzer, and interpreter

## Goal

Build a Go-based tool that can:
1. Parse PHP code into an AST
2. Analyze PHP code for types, issues, and insights
3. Execute PHP code directly (interpreter)
4. Transpile PHP code to equivalent Go code

## Current Status

- **Parser**: ✅ Working (lexer + recursive descent parser)
- **Analyzer**: ❌ Not started
- **Interpreter**: ❌ Not started
- **Transpiler**: ❌ Not started

## Tech Stack

- **Language**: Go 1.23
- **Linting**: golangci-lint
- **Testing**: Go's built-in testing
- **Pre-commit**: v6.0.0
- **CI**: GitHub Actions

## Key Files

```
main.go              - Entry point
cmd/root.go          - CLI commands
parser/parser.go    - PHP parser
internal/token/     - Lexer & tokens
internal/ast/      - AST node types
examples/           - PHP test examples
docs/               - Documentation
```

## Architecture

```
PHP Source → Lexer → Parser → AST → [Analyzer]
                       ↓                    ↓
                   [Interpreter]    [Transpiler] → Go Source
```

## Working Directory

```bash
/Users/azat/Projects/pet_projects/golang_stuff/phat9k
```

## Useful Commands

```bash
make test           # Run tests
make test-coverage # Generate coverage
make lint          # Run linter
make build         # Build binary
pre-commit run     # Run pre-commit hooks
go test ./...      # Run all tests
```