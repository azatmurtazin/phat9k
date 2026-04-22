# Feature: PHP Parser

## Description

The PHP Parser component transforms raw PHP source code into an Abstract Syntax Tree (AST) that can be processed by other components of Phat9k.

## Functionality

- Lexical analysis (tokenization) of PHP source code
- Syntax parsing into hierarchical AST nodes
- Support for PHP 7.x, 8.x and 8.5 syntax
- Error reporting with line and column information

## Implementation Details

Uses a recursive descent parser approach with token lookahead for efficient parsing. The parser handles:

- All control structures (if, else, elseif, switch, for, foreach, while, do-while)
- Function and method definitions
- Class definitions with inheritance
- Traits and interfaces
- Anonymous functions (closures)
- Expressions and operators
- Type declarations

## CLI Usage

```bash
phat9k parse input.php -o ast.json
```

Output is a JSON representation of the AST for debugging and tooling integration.