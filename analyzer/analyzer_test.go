package analyzer

import (
	"testing"

	"github.com/phat9k/parser"
)

func TestAnalyzer_FunctionDef(t *testing.T) {
	src := `<?php
function test() {
    return 1;
}
`
	p := parser.New(src)
	prog, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	a := New(prog)
	result, err := a.Analyze()
	if err != nil {
		t.Fatalf("analysis error: %v", err)
	}

	if result.Metrics.Functions != 1 {
		t.Errorf("expected 1 function, got %d", result.Metrics.Functions)
	}
}

func TestAnalyzer_ClassDef(t *testing.T) {
	src := `<?php
class Foo {
    public function bar() {
        return 1;
    }
}
`
	p := parser.New(src)
	prog, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	a := New(prog)
	result, err := a.Analyze()
	if err != nil {
		t.Fatalf("analysis error: %v", err)
	}

	if result.Metrics.Classes != 1 {
		t.Errorf("expected 1 class, got %d", result.Metrics.Classes)
	}
}

func TestAnalyzer_BothFunctionAndClass(t *testing.T) {
	src := `<?php
function foo() { return 1; }
class Bar { }
`
	p := parser.New(src)
	prog, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	a := New(prog)
	result, err := a.Analyze()
	if err != nil {
		t.Fatalf("analysis error: %v", err)
	}

	if result.Metrics.Functions != 1 {
		t.Errorf("expected 1 function, got %d", result.Metrics.Functions)
	}
	if result.Metrics.Classes != 1 {
		t.Errorf("expected 1 class, got %d", result.Metrics.Classes)
	}
}