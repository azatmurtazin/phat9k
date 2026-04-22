package parser

import (
	"fmt"

	"github.com/phat9k/internal/ast"
	"github.com/phat9k/internal/token"
)

type Parser struct {
	src    string
	tokens []token.Token
	pos    int
	depth  int
}

const maxDepth = 1000

func New(src string) *Parser {
	return &Parser{src: src}
}

func (p *Parser) Parse() (*ast.Program, error) {
	lexer := token.NewLexer(p.src)
	p.tokens = lexer.Tokenize()

	// Skip PHP opening tag
	if len(p.tokens) > 0 && p.tokens[0].Type == token.T_OPEN_TAG {
		p.pos++
	}

	stmts := []ast.Statement{}
	for p.pos < len(p.tokens) {
		if p.peek().Type == token.T_CLOSE_TAG {
			break
		}
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
	}

	return &ast.Program{Body: stmts}, nil
}

func (p *Parser) peek() token.Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return token.Token{Type: token.T_EOF}
}

func (p *Parser) advance() token.Token {
	tok := p.peek()
	p.pos++
	return tok
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	p.depth++
	if p.depth > 50 {
		return nil, fmt.Errorf("max parse depth exceeded: %d", p.depth)
	}
	defer func() { p.depth-- }()

	tok := p.peek()
	if tok.Type == token.T_EOF || tok.Type == token.T_CLOSE_TAG {
		return nil, nil
	}

	switch tok.Type {
	case token.T_ECHO:
		return p.parseEcho()
	case token.T_FUNCTION:
		return p.parseFunction()
	case token.T_CLASS:
		return p.parseClass()
	case token.T_IF:
		return p.parseIf()
	case token.T_FOR:
		return p.parseFor()
	case token.T_FOREACH:
		return p.parseForeach()
	case token.T_WHILE:
		return p.parseWhile()
	case token.T_RETURN:
		return p.parseReturn()
	case token.T_OPEN_TAG:
		p.advance()
		return p.parseStatement()
	default:
		if tok.Type == token.T_STRING && p.peekNext().Type == token.T_EQUAL {
			return p.parseAssignment()
		}
		// Skip unknown tokens to prevent infinite loop
		p.advance()
		if p.pos >= len(p.tokens) {
			return nil, nil
		}
		return p.parseStatement()
	}
}

func (p *Parser) peekNext() token.Token {
	if p.pos+1 < len(p.tokens) {
		return p.tokens[p.pos+1]
	}
	return token.Token{Type: token.T_EOF}
}

func (p *Parser) parseEcho() (*ast.EchoStatement, error) {
	p.advance() // skip echo

	exprs := []ast.Expression{}
	for p.peek().Type != token.T_SEMICOLON && p.peek().Type != token.T_CLOSE_TAG && p.peek().Type != token.T_EOF {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if expr != nil {
			exprs = append(exprs, expr)
		}
	}

	if p.peek().Type == token.T_SEMICOLON {
		p.advance()
	}

	return &ast.EchoStatement{Values: exprs}, nil
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	return p.parseBinaryExpr()
}

func (p *Parser) parseBinaryExpr() (ast.Expression, error) {
	left, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}

	for p.peek().Type == token.T_PLUS || p.peek().Type == token.T_MINUS ||
		p.peek().Type == token.T_MULT || p.peek().Type == token.T_DIV ||
		p.peek().Type == token.T_CONCAT {
		op := p.advance()
		right, err := p.parsePrimary()
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{
			Left:  left,
			Op:    op.Literal,
			Right: right,
		}
	}

	return left, nil
}

func (p *Parser) parsePrimary() (ast.Expression, error) {
	tok := p.peek()

	switch tok.Type {
	case token.T_STRING:
		p.advance()
		if p.peek().Type == token.T_LEFT_PAREN {
			return p.parseCall(tok)
		}
		return &ast.StringLiteral{Value: tok.Literal}, nil
	case token.T_LNUMBER, token.T_DNUMBER:
		p.advance()
		return &ast.NumberLiteral{Value: tok.Literal}, nil
	case token.T_CONSTANT_ENCAPSED_STRING:
		p.advance()
		return &ast.StringLiteral{Value: tok.Literal}, nil
	case token.T_VARIABLE:
		p.advance()
		return &ast.Variable{Name: tok.Literal}, nil
	case token.T_LEFT_PAREN:
		p.advance()
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		p.advance() // skip closing paren
		return expr, nil
	default:
		p.advance()
		return nil, nil
	}
}

func (p *Parser) parseCall(tok token.Token) (*ast.CallExpr, error) {
	p.advance() // skip function name
	p.advance() // skip (

	args := []ast.Expression{}
	for p.peek().Type != token.T_RIGHT_PAREN && p.peek().Type != token.T_EOF {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if expr != nil {
			args = append(args, expr)
		}
		if p.peek().Type == token.T_COMMA {
			p.advance()
		}
	}

	if p.peek().Type == token.T_RIGHT_PAREN {
		p.advance()
	}

	return &ast.CallExpr{Func: tok.Literal, Args: args}, nil
}

func (p *Parser) parseFunction() (*ast.FunctionDecl, error) {
	p.advance() // skip function

	nameTok := p.advance()
	name := nameTok.Literal

	if p.peek().Type == token.T_LEFT_PAREN {
		p.advance()
		for p.peek().Type != token.T_RIGHT_PAREN && p.peek().Type != token.T_EOF {
			p.advance()
			if p.peek().Type == token.T_COMMA {
				p.advance()
			}
		}
		if p.peek().Type == token.T_RIGHT_PAREN {
			p.advance()
		}
	}

	_ = p.parseBlock()

	return &ast.FunctionDecl{Name: name}, nil
}

func (p *Parser) parseClass() (*ast.ClassDecl, error) {
	p.advance() // skip class

	nameTok := p.advance()
	name := nameTok.Literal

	_ = p.parseBlock()

	return &ast.ClassDecl{Name: name}, nil
}

func (p *Parser) parseIf() (*ast.IfStatement, error) {
	p.advance() // skip if
	
	// Skip opening parenthesis if present
	if p.peek().Type == token.T_LEFT_PAREN {
		p.advance()
	}
	
	cond, _ := p.parseExpression()
	
	// Skip closing parenthesis if present
	if p.peek().Type == token.T_RIGHT_PAREN {
		p.advance()
	}
	
	then := p.parseBlock()

	var else_ []ast.Statement
	if p.peek().Type == token.T_ELSE {
		p.advance()
		else_ = p.parseBlock().Statements
	}

	return &ast.IfStatement{Condition: cond, Then: then, Else: else_}, nil
}

func (p *Parser) parseFor() (*ast.ForStatement, error) {
	p.advance() // skip for
	p.advance() // skip (
	_, _ = p.parseExpression()
	p.advance() // skip ;
	_, _ = p.parseExpression()
	p.advance() // skip ;
	_, _ = p.parseExpression()
	p.advance() // skip )
	body := p.parseBlock()

	return &ast.ForStatement{Body: body}, nil
}

func (p *Parser) parseForeach() (*ast.ForeachStatement, error) {
	p.advance() // skip foreach
	p.advance() // skip (
	_, _ = p.parseExpression()
	p.advance() // skip as
	_, _ = p.parseExpression()
	if p.peek().Type == token.T_DOUBLE_ARROW {
		p.advance()
		_, _ = p.parseExpression()
	}
	p.advance() // skip )
	body := p.parseBlock()

	return &ast.ForeachStatement{Body: body}, nil
}

func (p *Parser) parseWhile() (*ast.WhileStatement, error) {
	p.advance() // skip while
	p.advance() // skip (
	_, _ = p.parseExpression()
	p.advance() // skip )
	body := p.parseBlock()

	return &ast.WhileStatement{Body: body}, nil
}

func (p *Parser) parseReturn() (*ast.ReturnStatement, error) {
	p.advance() // skip return

	var val ast.Expression
	if p.peek().Type != token.T_SEMICOLON {
		val, _ = p.parseExpression()
	}

	if p.peek().Type == token.T_SEMICOLON {
		p.advance()
	}

	return &ast.ReturnStatement{Value: val}, nil
}

func (p *Parser) parseAssignment() (*ast.Assignment, error) {
	name := p.advance()
	p.advance() // skip =
	val, _ := p.parseExpression()
	if p.peek().Type == token.T_SEMICOLON {
		p.advance()
	}

	return &ast.Assignment{Name: name.Literal, Value: val}, nil
}

func (p *Parser) parseBlock() *ast.Block {
	if p.peek().Type != token.T_LEFT_BRACE {
		return &ast.Block{}
	}
	p.advance() // skip {

	stmts := []ast.Statement{}
	for p.peek().Type != token.T_RIGHT_BRACE && p.peek().Type != token.T_EOF {
		stmt, _ := p.parseStatement()
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
	}

	if p.peek().Type == token.T_RIGHT_BRACE {
		p.advance()
	}

	return &ast.Block{Statements: stmts}
}
