package parser

import (
	"os"
	"testing"
)

func TestIntegration_ParseSimplePHP(t *testing.T) {
	src := `<?php
$a = 1;
$b = 2;
echo $a + $b;
`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if ast == nil {
		t.Fatal("expected AST, got nil")
	}
}

func TestIntegration_ParseFunction(t *testing.T) {
	src := `<?php
function greet($name) {
    return "Hello, " . $name;
}
echo greet("World");
`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if ast == nil {
		t.Fatal("expected AST, got nil")
	}
}

func TestIntegration_ParseClass(t *testing.T) {
	src := `<?php
class User {
    public $name;

    public function __construct($name) {
        $this->name = $name;
    }

    public function getName() {
        return $this->name;
    }
}
`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if ast == nil {
		t.Fatal("expected AST, got nil")
	}
}

func TestIntegration_ParseFile(t *testing.T) {
	src, err := os.ReadFile("testdata/parser/simple.php")
	if err != nil {
		t.Skip("test file not found")
	}

	p := New(string(src))
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if ast == nil {
		t.Fatal("expected AST, got nil")
	}
}