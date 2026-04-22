package token

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type Lexer struct {
	src   string
	pos   int
	line  int
	column int
}

func NewLexer(src string) *Lexer {
	src = strings.TrimLeft(src, "\uFEFF")
	if strings.HasPrefix(src, "<?php") {
		src = strings.TrimPrefix(src, "<?php")
		src = strings.TrimLeft(src, " \t\r\n")
	}
	return &Lexer{src: src, line: 1, column: 1}
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token

	for l.pos < len(l.src) {
		tok := l.scanToken()
		tokens = append(tokens, tok)
	}

	tokens = append(tokens, Token{Type: T_EOF, Line: l.line, Column: l.column})

	return tokens
}

func (l *Lexer) scanToken() Token {
	if l.pos >= len(l.src) {
		return Token{Type: T_EOF, Line: l.line, Column: l.column}
	}

	c := l.src[l.pos]

	if c == '\n' {
		l.pos++
		l.line++
		l.column = 1
		return l.scanToken()
	}

	if unicode.IsSpace(rune(c)) {
		l.pos++
		l.column++
		return l.scanToken()
	}

	if c == '=' {
		l.pos++
		l.column++
		if l.pos < len(l.src) && l.src[l.pos] == '>' {
			l.pos++
			l.column++
			return Token{Type: T_DOUBLE_ARROW, Literal: "=>", Line: l.line, Column: l.column - 2}
		}
		return Token{Type: T_EQUAL, Literal: "=", Line: l.line, Column: l.column - 1}
	}

	if c == '"' || c == '\'' {
		return l.scanString(c)
	}

	if unicode.IsDigit(rune(c)) {
		return l.scanNumber()
	}

	if unicode.IsLetter(rune(c)) || c == '_' || c == '$' {
		return l.scanIdent()
	}

	return l.scanPunct()
}

func (l *Lexer) scanString(quote byte) Token {
	start := l.pos
	l.pos++
	l.column++

	for l.pos < len(l.src) && l.src[l.pos] != quote {
		if l.src[l.pos] == '\\' && l.pos+1 < len(l.src) {
			l.pos += 2
			l.column += 2
			continue
		}
		if l.src[l.pos] == '\n' {
			l.pos++
			l.line++
			l.column = 1
			continue
		}
		l.pos++
		l.column++
	}

	if l.pos < len(l.src) {
		l.pos++
		l.column++
	}

	value := l.src[start:l.pos]

	if quote == '"' {
		return Token{Type: T_CONSTANT_ENCAPSED_STRING, Literal: value, Line: l.line, Column: l.column - len(value)}
	}
	return Token{Type: T_CONSTANT_ENCAPSED_STRING, Literal: value, Line: l.line, Column: l.column - len(value)}
}

func (l *Lexer) scanNumber() Token {
	start := l.pos
	hasDot := false

	for l.pos < len(l.src) {
		c := l.src[l.pos]
		if unicode.IsDigit(rune(c)) {
			l.pos++
			l.column++
			continue
		}
		if c == '.' && !hasDot && l.pos+1 < len(l.src) && unicode.IsDigit(rune(l.src[l.pos+1])) {
			hasDot = true
			l.pos++
			l.column++
			continue
		}
		break
	}

	value := l.src[start:l.pos]
	l.column += len(value)

	if hasDot {
		return Token{Type: T_DNUMBER, Literal: value, Line: l.line, Column: l.column - len(value)}
	}
	return Token{Type: T_LNUMBER, Literal: value, Line: l.line, Column: l.column - len(value)}
}

func (l *Lexer) scanIdent() Token {
	start := l.pos

	for l.pos < len(l.src) {
		c := l.src[l.pos]
		if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '_' {
			l.pos++
			l.column++
			continue
		}
		break
	}

	value := l.src[start:l.pos]

	if kw, ok := keywords[value]; ok {
		return Token{Type: kw, Literal: value, Line: l.line, Column: l.column - len(value)}
	}

	if len(value) > 0 && value[0] == '$' && len(value) > 1 {
		return Token{Type: T_VARIABLE, Literal: value, Line: l.line, Column: l.column - len(value)}
	}

	return Token{Type: T_STRING, Literal: value, Line: l.line, Column: l.column - len(value)}
}

func (l *Lexer) scanPunct() Token {
	startCol := l.column
	c := l.src[l.pos]
	l.pos++
	l.column++

	switch c {
	case '(':
		return Token{Type: T_LEFT_PAREN, Literal: "(", Line: l.line, Column: startCol}
	case ')':
		return Token{Type: T_RIGHT_PAREN, Literal: ")", Line: l.line, Column: startCol}
	case '[':
		return Token{Type: T_LEFT_BRACKET, Literal: "[", Line: l.line, Column: startCol}
	case ']':
		return Token{Type: T_RIGHT_BRACKET, Literal: "]", Line: l.line, Column: startCol}
	case '{':
		return Token{Type: T_LEFT_BRACE, Literal: "{", Line: l.line, Column: startCol}
	case '}':
		return Token{Type: T_RIGHT_BRACE, Literal: "}", Line: l.line, Column: startCol}
	case ';':
		return Token{Type: T_SEMICOLON, Literal: ";", Line: l.line, Column: startCol}
	case ',':
		return Token{Type: T_COMMA, Literal: ",", Line: l.line, Column: startCol}
	case ':':
		return Token{Type: T_COLON, Literal: ":", Line: l.line, Column: startCol}
	case '.':
		return Token{Type: T_CONCAT, Literal: ".", Line: l.line, Column: startCol}
	case '+':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_PLUS_EQUAL, Literal: "+=", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '+' {
			l.pos++
			l.column++
			return Token{Type: T_INC, Literal: "++", Line: l.line, Column: startCol}
		}
		return Token{Type: T_PLUS, Literal: "+", Line: l.line, Column: startCol}
	case '-':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_MINUS_EQUAL, Literal: "-=", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '-' {
			l.pos++
			l.column++
			return Token{Type: T_DEC, Literal: "--", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '>' {
			l.pos++
			l.column++
			return Token{Type: T_NS_CODEL, Literal: "->", Line: l.line, Column: startCol}
		}
		return Token{Type: T_MINUS, Literal: "-", Line: l.line, Column: startCol}
	case '*':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_MULT_EQUAL, Literal: "*=", Line: l.line, Column: startCol}
		}
		return Token{Type: T_MULT, Literal: "*", Line: l.line, Column: startCol}
	case '/':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_DIV_EQUAL, Literal: "/=", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '/' {
			return l.scanComment()
		}
		if l.pos < len(l.src) && l.src[l.pos] == '*' {
			return l.scanBlockComment()
		}
		return Token{Type: T_DIV, Literal: "/", Line: l.line, Column: startCol}
	case '%':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_MOD_EQUAL, Literal: "%=", Line: l.line, Column: startCol}
		}
		return Token{Type: T_MOD, Literal: "%", Line: l.line, Column: startCol}
	case '<':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_SMALLER_OR_EQUAL, Literal: "<=", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '?' {
			l.pos++
			l.column++
			return Token{Type: T_OPEN_TAG, Literal: "<?", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '<' {
			l.pos++
			l.column++
			if l.pos < len(l.src) && l.src[l.pos] == '=' {
				l.pos++
				l.column++
				return Token{Type: T_SL_EQUAL, Literal: "<<=", Line: l.line, Column: startCol}
			}
			return Token{Type: T_SL, Literal: "<<", Line: l.line, Column: startCol}
		}
		return Token{Type: T_SMALLER, Literal: "<", Line: l.line, Column: startCol}
	case '>':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_GREATER_OR_EQUAL, Literal: ">=", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '>' {
			l.pos++
			l.column++
			if l.pos < len(l.src) && l.src[l.pos] == '=' {
				l.pos++
				l.column++
				return Token{Type: T_SR_EQUAL, Literal: ">>=", Line: l.line, Column: startCol}
			}
			return Token{Type: T_SR, Literal: ">>", Line: l.line, Column: startCol}
		}
		return Token{Type: T_GREATER, Literal: ">", Line: l.line, Column: startCol}
	case '!':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			if l.pos < len(l.src) && l.src[l.pos] == '=' {
				l.pos++
				l.column++
				return Token{Type: T_IS_NOT_IDENTICAL, Literal: "!==", Line: l.line, Column: startCol}
			}
			return Token{Type: T_IS_NOT_EQUAL, Literal: "!=", Line: l.line, Column: startCol}
		}
		return Token{Type: T_BOOLEAN_NOT, Literal: "!", Line: l.line, Column: startCol}
	case '&':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_BITWISE_AND_EQUAL, Literal: "&=", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '&' {
			l.pos++
			l.column++
			return Token{Type: T_BOOLEAN_AND, Literal: "&&", Line: l.line, Column: startCol}
		}
		return Token{Type: T_AMPERSAND, Literal: "&", Line: l.line, Column: startCol}
	case '|':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_BITWISE_OR_EQUAL, Literal: "|=", Line: l.line, Column: startCol}
		}
		if l.pos < len(l.src) && l.src[l.pos] == '|' {
			l.pos++
			l.column++
			return Token{Type: T_BOOLEAN_OR, Literal: "||", Line: l.line, Column: startCol}
		}
		return Token{Type: T_BITWISE_AND, Literal: "|", Line: l.line, Column: startCol}
	case '^':
		if l.pos < len(l.src) && l.src[l.pos] == '=' {
			l.pos++
			l.column++
			return Token{Type: T_BITWISE_XOR_EQUAL, Literal: "^=", Line: l.line, Column: startCol}
		}
		return Token{Type: T_BITWISE_AND, Literal: "^", Line: l.line, Column: startCol}
	case '?':
		if l.pos < len(l.src) && l.src[l.pos] == '?' {
			l.pos++
			l.column++
			if l.pos < len(l.src) && l.src[l.pos] == '=' {
				l.pos++
				l.column++
				return Token{Type: T_COALESCE_EQUAL, Literal: "??=", Line: l.line, Column: startCol}
			}
			return Token{Type: T_COALESCE, Literal: "??", Line: l.line, Column: startCol}
		}
		return Token{Type: T_QUESTION, Literal: "?", Line: l.line, Column: startCol}
	case '@':
		return Token{Type: T_AT, Literal: "@", Line: l.line, Column: startCol}
	case '~':
		return Token{Type: T_BITWISE_AND, Literal: "~", Line: l.line, Column: startCol}
	case '$':
		return l.scanVariable()
	}

	return Token{Type: T_STRING, Literal: string(c), Line: l.line, Column: startCol}
}

func (l *Lexer) scanVariable() Token {
	start := l.pos - 1
	startCol := l.column - 1

	for l.pos < len(l.src) {
		c := l.src[l.pos]
		if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '_' {
			l.pos++
			l.column++
			continue
		}
		break
	}

	value := l.src[start:l.pos]
	if value == "$" {
		return Token{Type: T_VARIABLE, Literal: "$", Line: l.line, Column: startCol}
	}
	return Token{Type: T_VARIABLE, Literal: value, Line: l.line, Column: startCol}
}

func (l *Lexer) scanComment() Token {
	start := l.pos - 1
	startCol := l.column - 1

	for l.pos < len(l.src) && l.src[l.pos] != '\n' {
		l.pos++
		l.column++
	}

	return Token{Type: T_STRING, Literal: l.src[start:l.pos], Line: l.line, Column: startCol}
}

func (l *Lexer) scanBlockComment() Token {
	l.pos += 2
	l.column += 2

	for l.pos+1 < len(l.src) {
		if l.src[l.pos] == '*' && l.src[l.pos+1] == '/' {
			l.pos += 2
			l.column += 2
			break
		}
		if l.src[l.pos] == '\n' {
			l.line++
			l.column = 1
		}
		l.pos++
		l.column++
	}

	return Token{Type: T_STRING, Literal: "comment", Line: l.line, Column: 1}
}

var keywords = map[string]TokenType{
	"echo":       T_ECHO,
	"print":      T_PRINT,
	"function":   T_FUNCTION,
	"class":      T_CLASS,
	"interface":  T_INTERFACE,
	"trait":      T_TRAIT,
	"extends":   T_EXTENDS,
	"implements": T_IMPLEMENTS,
	"new":        T_NEW,
	"if":         T_IF,
	"elseif":     T_ELSEIF,
	"else":       T_ELSE,
	"switch":     T_SWITCH,
	"case":       T_CASE,
	"default":   T_DEFAULT,
	"for":        T_FOR,
	"foreach":    T_FOREACH,
	"while":     T_WHILE,
	"do":         T_DO,
	"return":     T_RETURN,
	"break":      T_BREAK,
	"continue":   T_CONTINUE,
	"exit":       T_EXIT,
	"die":        T_DIE,
	"try":        T_TRY,
	"catch":      T_CATCH,
	"throw":      T_THROW,
	"namespace": T_NAMESPACE,
	"use":        T_USE,
	"public":    T_PUBLIC,
	"private":    T_PRIVATE,
	"protected": T_PROTECTED,
	"static":     T_STATIC,
	"final":      T_FINAL,
	"abstract":  T_ABSTRACT,
	"const":      T_CONST,
	"var":        T_VAR,
	"array":      T_ARRAY,
	"list":       T_LIST,
	"global":     T_GLOBAL,
	"unset":      T_UNSET,
	"empty":      T_EMPTY,
	"isset":      T_ISSET,
	"include":   T_INCLUDE,
	"include_once": T_INCLUDE_ONCE,
	"require":   T_REQUIRE,
	"require_once": T_REQUIRE_ONCE,
	"true":       T_TRUE,
	"false":      T_FALSE,
	"null":       T_NULL,
	"and":        T_LOGICAL_AND,
	"or":         T_LOGICAL_OR,
	"xor":        T_LOGICAL_XOR,
	"instanceof": T_INSTANCEOF,
}

var _ = fmt.Sprintf
var _ = regexp.MustCompile
