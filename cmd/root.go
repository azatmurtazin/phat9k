package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/phat9k/analyzer"
	"github.com/phat9k/interpreter"
	"github.com/phat9k/parser"
)

func Run(args []string) error {
	if len(args) < 2 {
		return usage()
	}

	switch args[1] {
	case "parse":
		return runParse(args[2:])
	case "analyze":
		return runAnalyze(args[2:])
	case "run":
		return runRun(args[2:])
	case "transpile":
		return runTranspile(args[2:])
	case "help", "--help", "-h":
		return usage()
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func usage() error {
	fmt.Println(`Phat9k - PHP code analyzer, interpreter and transpiler

Usage:
  phat9k <command> [options]

Commands:
  parse      Parse PHP code into AST
  analyze   Analyze PHP code for types and issues
  run       Execute PHP code
  transpile Convert PHP code to Go

Use "phat9k <command> -h" for more information about a command.`)
	return nil
}

func runParse(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: phat9k parse <file>")
	}

	src, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	p := parser.New(string(src))
	ast, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	fmt.Println(ast.String())
	return nil
}

func runAnalyze(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: phat9k analyze <file>")
	}

	src, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	p := parser.New(string(src))
	ast, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	a := analyzer.New(ast)
	result, err := a.Analyze()
	if err != nil {
		return fmt.Errorf("analyzing: %w", err)
	}

	output := "text"
	for _, arg := range args[1:] {
		if arg == "-o" || arg == "--output" {
			output = "json"
		}
	}

	if output == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(result)
	} else {
		fmt.Println(analyzer.FormatResult(result))
	}
	return nil
}

func runRun(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: phat9k run <file>")
	}

	src, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	p := parser.New(string(src))
	ast, err := p.Parse()
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	r := interpreter.New(ast)
	result := r.Execute()
	if result.Error != nil {
		return fmt.Errorf("runtime error: %w", result.Error)
	}

	fmt.Print(result.Output)
	return nil
}

func runTranspile(args []string) error {
	return fmt.Errorf("not implemented")
}
