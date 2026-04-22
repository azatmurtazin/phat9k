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
- **Analyzer**: ✅ Working (type inference, metrics, error detection)
- **Interpreter**: ✅ Working (variables, expressions, basic functions)
- **Transpiler**: ✅ Working (PHP to Go code generation)
- **CLI**: ✅ All commands (parse, analyze, run, transpile)

## Tech Stack

- **Language**: Go 1.23
- **Linting**: golangci-lint
- **Testing**: Go's built-in testing
- **Pre-commit**: v6.0.0
- **CI**: GitHub Actions

## Key Files

```
cmd/phat9k/main.go   - CLI entry point (binary: ./bin/phat9k)
cmd/root.go        - CLI commands
parser/parser.go   - PHP parser
analyzer/         - PHP code analyzer
interpreter/      - PHP interpreter
transpiler/        - PHP to Go transpiler
internal/token/    - Lexer & tokens
internal/ast/      - AST node types
examples/         - PHP test examples
tests.yml         - Test results matrix
docs/              - Documentation
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