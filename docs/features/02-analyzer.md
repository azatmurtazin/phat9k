# Feature: PHP Analyzer

## Description

The PHP Analyzer performs static analysis on PHP code to understand types, detect potential issues, and generate insights for transpilation.

## Functionality

- Type inference from variable usage and return values
- Detection of undefined variables and functions
- Dead code detection
- Security vulnerability scanning
- Code complexity analysis
- Call graph generation

## Implementation Details

Walks the AST to collect:

- Variable definitions and usages
- Function and method signatures
- Class hierarchies
- Type annotations (from PHPDoc, PHP 8 attributes, and PHP 8.5 annotations)
- Dynamic code patterns

## CLI Usage

```bash
phat9k analyze input.php -o analysis.json
```

Output includes type information, warnings, errors, and metrics.