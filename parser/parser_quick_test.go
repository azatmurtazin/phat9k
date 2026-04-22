package parser

import (
	"testing"
)

func TestParser_SimpleEcho(t *testing.T) {
	src := `echo "hello";`
	p := New(src)
	_, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
}

func TestParser_Assignment(t *testing.T) {
	src := `$x = 1;`
	p := New(src)
	_, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
}

func TestParser_Math(t *testing.T) {
	src := ` 1 + 2; `
	p := New(src)
	_, err := p.Parse()
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
}