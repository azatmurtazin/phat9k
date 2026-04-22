# Feature: PHP to Go Transpiler

## Description

The Transpiler converts PHP code to equivalent Go code, enabling migration from PHP to Go.

## Functionality

- Translate PHP syntax to Go syntax
- Map PHP types to Go types
- Convert PHP standard library to Go equivalents
- Support for PHP 8.5 features
- Preserve code structure and logic
- Generate Go module structure

## Implementation Details

Transforms the AST by:

- Mapping control structures (if/else/switch to Go's switch with tags)
- Converting functions to Go functions
- Translating classes to Go structs with methods
- Handling dynamic features with reflectivity or code generation
- Mapping built-in functions to Go packages (e.g., strings, fmt, http)

### Type Mapping

| PHP Type | Go Type |
|----------|---------|
| int | int |
| float | float64 |
| string | string |
| bool | bool |
| array | map[string]interface{} or []interface{} |
| null | nil |
| callable | func(...interface{}) interface{} |

### Function Mapping

- `echo` / `print` -> `fmt.Print`
- `strlen` -> `len`
- `strpos` -> `strings.Index`
- `array_map` -> custom mapping function
- etc.

## CLI Usage

```bash
phat9k transpile input.php -o output.go
```

Generates Go source file from PHP input.
