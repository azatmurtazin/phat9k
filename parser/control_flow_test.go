package parser

import (
	"testing"
)

func TestParser_IfStatement(t *testing.T) {
	src := `<?php if ($x) { echo 1; }`
	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if ast == nil {
		t.Fatal("nil AST")
	}
}

func TestParser_ForStatement(t *testing.T) {
	src := `<?php for ($i = 0; $i < 10; $i++) { }`
	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if ast == nil {
		t.Fatal("nil AST")
	}
}