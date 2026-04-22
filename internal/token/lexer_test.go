package token

import (
	"testing"
)

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		name  string
		src   string
		check func([]Token)
	}{
		{
			name: "empty string",
			src:  "",
			check: func(ts []Token) {
				if len(ts) != 1 {
					t.Errorf("expected 1 token, got %d", len(ts))
				}
				if ts[0].Type != T_EOF {
					t.Errorf("expected EOF, got %v", ts[0].Type)
				}
			},
		},
		{
			name: "echo statement",
			src:  "echo 'hello';",
			check: func(ts []Token) {
				if len(ts) < 2 {
					t.Fatalf("expected at least 2 tokens, got %d", len(ts))
				}
				if ts[0].Type != T_ECHO {
					t.Errorf("expected echo, got %v", ts[0].Type)
				}
			},
		},
		{
			name: "number",
			src:  "42",
			check: func(ts []Token) {
				if len(ts) < 1 {
					t.Fatalf("expected at least 1 token, got %d", len(ts))
				}
				if ts[0].Type != T_LNUMBER {
					t.Errorf("expected LNUMBER, got %v", ts[0].Type)
				}
				if ts[0].Literal != "42" {
					t.Errorf("expected '42', got %s", ts[0].Literal)
				}
			},
		},
		{
			name: "variable",
			src:  "$foo",
			check: func(ts []Token) {
				if len(ts) < 1 {
					t.Fatalf("expected at least 1 token, got %d", len(ts))
				}
				if ts[0].Type != T_VARIABLE {
					t.Errorf("expected VARIABLE, got %v", ts[0].Type)
				}
				if ts[0].Literal != "$foo" {
					t.Errorf("expected '$foo', got %s", ts[0].Literal)
				}
			},
		},
		{
			name: "addition",
			src:  "1 + 2",
			check: func(ts []Token) {
				if len(ts) < 3 {
					t.Fatalf("expected at least 3 tokens, got %d", len(ts))
				}
				if ts[1].Type != T_PLUS {
					t.Errorf("expected PLUS, got %v", ts[1].Type)
				}
			},
		},
		{
			name: "function call",
			src:  "foo()",
			check: func(ts []Token) {
				if len(ts) < 3 {
					t.Fatalf("expected at least 3 tokens, got %d", len(ts))
				}
				if ts[0].Type != T_STRING {
					t.Errorf("expected STRING, got %v", ts[0].Type)
				}
				if ts[0].Literal != "foo" {
					t.Errorf("expected 'foo', got %s", ts[0].Literal)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.src)
			ts := lexer.Tokenize()
			tt.check(ts)
		})
	}
}

func TestLexer_LineColumn(t *testing.T) {
	src := "echo 'test';\n$var = 1;"
	lexer := NewLexer(src)
	ts := lexer.Tokenize()

	var last Token
	for _, tok := range ts {
		last = tok
	}
	if last.Line < 1 {
		t.Errorf("expected line >= 1, got %d", last.Line)
	}
}

func TestLexer_Keywords(t *testing.T) {
	keywords := []string{
		"echo", "print", "function", "class", "interface", "trait",
		"if", "else", "elseif", "switch", "case", "default",
		"for", "foreach", "while", "do", "return", "break", "continue",
		"new", "instanceof", "extends", "implements",
		"public", "private", "protected", "static", "final", "abstract",
		"const", "var", "array", "true", "false", "null",
	}

	for _, kw := range keywords {
		src := kw
		lexer := NewLexer(src)
		ts := lexer.Tokenize()
		if len(ts) < 1 {
			t.Errorf("expected at least 1 token for keyword %s", kw)
			continue
		}
		if ts[0].Type == T_STRING {
			t.Errorf("expected keyword token for %s, got STRING", kw)
		}
	}
}
