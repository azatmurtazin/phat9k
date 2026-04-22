package transpiler

import (
	"fmt"
	"strings"

	"github.com/azatmurtazin/phat9k/internal/ast"
)

type Transpiler struct {
	program  *ast.Program
	output   strings.Builder
	indent   int
	functions map[string]bool
	imports  map[string]bool
}

func New(program *ast.Program) *Transpiler {
	return &Transpiler{
		program:  program,
		output:   strings.Builder{},
		functions: make(map[string]bool),
		imports:  make(map[string]bool),
	}
}

func (t *Transpiler) Transpile() (string, error) {
	t.output.WriteString("package main\n\n")
	t.output.WriteString("import (\n")

	t.output.WriteString("\t\"fmt\"\n")
	t.output.WriteString(")\n\n")

	t.output.WriteString("func main() {\n")
	t.indent++

	for _, stmt := range t.program.Body {
		t.transpileStmt(stmt)
	}

	t.indent--
	t.output.WriteString("}\n")

	for name := range t.functions {
		t.output.WriteString("\n")
		t.output.WriteString(fmt.Sprintf("func %s() {\n", name))
		t.output.WriteString("\t// TODO: implement\n")
		t.output.WriteString("}\n")
	}

	return t.output.String(), nil
}

func (t *Transpiler) transpileStmt(node ast.Node) {
	if node == nil {
		return
	}

	switch n := node.(type) {
	case *ast.EchoStatement:
		for _, expr := range n.Values {
			val := t.transpileExpr(expr)
			t.output.WriteString(fmt.Sprintf("fmt.Println(%s)\n", val))
		}

	case *ast.Assignment:
		name := "$" + n.Name
		val := t.transpileExpr(n.Value)
		t.output.WriteString(fmt.Sprintf("%s := %s\n", name, val))

	case *ast.Block:
		for _, stmt := range n.Statements {
			t.transpileStmt(stmt)
		}

	case *ast.FunctionDecl:
		t.functions[n.Name] = true
		t.output.WriteString(fmt.Sprintf("func %s() {\n", n.Name))
		t.indent++
		if n.Body != nil {
			for _, stmt := range n.Body.Statements {
				t.transpileStmt(stmt)
			}
		}
		t.indent--
		t.output.WriteString("}\n")

	case *ast.ClassDecl:
		t.output.WriteString(fmt.Sprintf("type %s struct {\n", n.Name))
		t.indent++
		t.output.WriteString("}\n")
		t.indent--
	}

	t.output.WriteString("\n")
}

func (t *Transpiler) transpileExpr(node ast.Node) string {
	if node == nil {
		return "nil"
	}

	switch n := node.(type) {
	case *ast.NumberLiteral:
		return n.Value

	case *ast.StringLiteral:
		s := strings.Trim(n.Value, `"'`)
		return fmt.Sprintf(`"%s"`, s)

	case *ast.Variable:
		return "$" + n.Name

	case *ast.BinaryExpr:
		left := t.transpileExpr(n.Left)
		right := t.transpileExpr(n.Right)
		op := n.Op

		if op == "." {
			op = "+"
			t.imports["strings"] = true
			return fmt.Sprintf("fmt.Sprintf(\"%%s%%s\", %s, %s)", left, right)
		}

		switch n.Op {
		case "+", "-", "*", "/", "%":
			return fmt.Sprintf("(%s %s %s)", left, op, right)
		}
		return fmt.Sprintf("(%s %s %s)", left, op, right)

	case *ast.CallExpr:
		args := ""
		if n.Args != nil {
			var a []string
			for _, arg := range n.Args {
				a = append(a, t.transpileExpr(arg))
			}
			args = strings.Join(a, ", ")
		}
		return fmt.Sprintf("%s(%s)", n.Func, args)
	}

	return "nil"
}