package parser

import (
	"testing"
)

func TestParser_ParseEcho(t *testing.T) {
	src := `echo "hello", "world";`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf(" parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseAssignment(t *testing.T) {
	src := `$x = 1;`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseFunction(t *testing.T) {
	src := `function foo() { return 1; }`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseClass(t *testing.T) {
	src := `class Foo { public $bar; }`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseIf(t *testing.T) {
	src := `if ($x) { echo "yes"; }`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseFor(t *testing.T) {
	src := `for ($i = 0; $i < 10; $i++) { echo $i; }`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseForeach(t *testing.T) {
	src := `foreach ($arr as $v) { echo $v; }`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseWhile(t *testing.T) {
	src := `while ($x) { $x = false; }`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseReturn(t *testing.T) {
	src := `function test() { return 42; }`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(ast.Body))
	}
}

func TestParser_ParseMultiple(t *testing.T) {
	src := `
$a = 1;
$b = 2;
echo $a + $b;
`

	p := New(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if len(ast.Body) < 2 {
		t.Fatalf("expected at least 2 statements, got %d", len(ast.Body))
	}
}