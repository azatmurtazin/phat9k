# Feature: PHP Interpreter

## Description

The PHP Interpreter executes PHP code directly without compilation, useful for testing and quick validation.

## Functionality

- Execute PHP code in a sandboxed environment
- Support for core PHP functions and extensions
- Capture output and return values
- Debugging console mode

## Implementation Details

Interprets the AST directly with:

- Scope management for variables
- Built-in function registry
- Exception handling
- Memory-safe execution limits

## CLI Usage

```bash
phat9k run input.php
```

Executes the PHP file and outputs result to stdout.