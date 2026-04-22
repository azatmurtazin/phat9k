package main

import (
	"testing"

	"github.com/phat9k/analyzer"
	"github.com/phat9k/parser"
)

func TestExample_Echo(t *testing.T) {
	src := `echo "hello";`
	p := parser.New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if ast == nil {
		t.Fatal("expected AST")
	}
}

func TestExample_Analyzer(t *testing.T) {
	src := `<?php function foo() { return 1; }`
	p := parser.New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	a := analyzer.New(ast)
	result, err := a.Analyze()
	if err != nil {
		t.Fatalf("analyze error: %v", err)
	}

	if result.Metrics.Functions != 1 {
		t.Errorf("expected 1 function, got %d", result.Metrics.Functions)
	}
}