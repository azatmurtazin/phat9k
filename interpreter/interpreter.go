package interpreter

import (
	"fmt"
	"strings"

	"github.com/phat9k/internal/ast"
)

type Value interface{}

type Runtime struct {
	program  *ast.Program
	scope    *Scope
	output   strings.Builder
}

type Scope struct {
	parent *Scope
	vars   map[string]Value
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		vars:   make(map[string]Value),
	}
}

func (s *Scope) Get(name string) Value {
	if v, ok := s.vars[name]; ok {
		return v
	}
	if s.parent != nil {
		return s.parent.Get(name)
	}
	return nil
}

func (s *Scope) Set(name string, val Value) {
	s.vars[name] = val
}

type Result struct {
	Value  Value
	Output string
	Error error
}

func New(program *ast.Program) *Runtime {
	return &Runtime{
		program: program,
		scope:   NewScope(nil),
		output:   strings.Builder{},
	}
}

func (r *Runtime) Execute() *Result {
	err := r.run()
	if err != nil {
		return &Result{Error: err}
	}
	return &Result{Output: r.output.String()}
}

func (r *Runtime) run() error {
	for _, stmt := range r.program.Body {
		err := r.exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Runtime) exec(node ast.Node) error {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.Assignment:
		val := r.eval(n.Value)
		r.scope.Set(n.Name, val)

	case *ast.EchoStatement:
		for _, expr := range n.Values {
			val := r.eval(expr)
			r.output.WriteString(r.str(val))
		}

	case *ast.ReturnStatement:
		return nil

	case *ast.Block:
		for _, stmt := range n.Statements {
			err := r.exec(stmt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Runtime) eval(node ast.Node) Value {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.NumberLiteral:
		return r.parseNum(n.Value)

	case *ast.StringLiteral:
		return r.unquote(n.Value)

	case *ast.Variable:
		return r.scope.Get(n.Name)

	case *ast.BinaryExpr:
		left := r.eval(n.Left)
		right := r.eval(n.Right)
		return r.binOp(n.Op, left, right)
	}
	return nil
}

func (r *Runtime) str(v Value) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("%d", x)
	case string:
		return x
	default:
		if v == nil {
			return ""
		}
		return fmt.Sprintf("%v", v)
	}
}

func (r *Runtime) unquote(s string) string {
	s = strings.TrimPrefix(s, "'")
	s = strings.TrimPrefix(s, `"`)
	s = strings.TrimSuffix(s, "'")
	s = strings.TrimSuffix(s, `"`)
	return s
}

func (r *Runtime) parseNum(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func (r *Runtime) binOp(op string, a, b Value) Value {
	ai, aok := a.(int)
	bi, bok := b.(int)
	if aok && bok {
		switch op {
		case "+":
			return ai + bi
		case "-":
			return ai - bi
		case "*":
			return ai * bi
		case "/":
			if bi != 0 {
				return ai / bi
			}
		case "%":
			if bi != 0 {
				return ai % bi
			}
		}
	}
	as, aok := a.(string)
	bs, bok := b.(string)
	if aok && bok && op == "." {
		return as + bs
	}
	return nil
}