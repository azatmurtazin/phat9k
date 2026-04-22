package ast

import (
	"fmt"
	"strings"
)

type Node interface {
	NodeType() string
	String() string
}

type Program struct {
	Body []Statement
}

func (p *Program) NodeType() string { return "Program" }
func (p *Program) String() string {
	var b strings.Builder
	for _, s := range p.Body {
		b.WriteString(s.String())
		b.WriteString("\n")
	}
	return b.String()
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Literal struct {
	Value string
}

func (l *Literal) expressionNode()  {}
func (l *Literal) NodeType() string { return "Literal" }
func (l *Literal) String() string   { return l.Value }

type NumberLiteral struct {
	Value string
}

func (n *NumberLiteral) expressionNode()  {}
func (n *NumberLiteral) NodeType() string { return "NumberLiteral" }
func (n *NumberLiteral) String() string   { return n.Value }

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) expressionNode()  {}
func (s *StringLiteral) NodeType() string { return "StringLiteral" }
func (s *StringLiteral) String() string   { return s.Value }

type Variable struct {
	Name string
}

func (v *Variable) expressionNode()  {}
func (v *Variable) NodeType() string { return "Variable" }
func (v *Variable) String() string   { return v.Name }

type BinaryExpr struct {
	Left  Expression
	Op    string
	Right Expression
}

func (b *BinaryExpr) expressionNode()  {}
func (b *BinaryExpr) NodeType() string { return "BinaryExpr" }
func (b *BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left, b.Op, b.Right)
}

type UnaryExpr struct {
	Op   string
	Expr Expression
}

func (u *UnaryExpr) NodeType() string { return "UnaryExpr" }
func (u *UnaryExpr) String() string {
	return fmt.Sprintf("(%s%s)", u.Op, u.Expr)
}

type CallExpr struct {
	Func string
	Args []Expression
}

func (c *CallExpr) expressionNode()  {}
func (c *CallExpr) NodeType() string { return "CallExpr" }
func (c *CallExpr) String() string {
	var a []string
	for _, arg := range c.Args {
		a = append(a, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.Func, strings.Join(a, ", "))
}

type ArrayExpr struct {
	Elements []ArrayElement
}

func (a *ArrayExpr) NodeType() string { return "ArrayExpr" }
func (a *ArrayExpr) String() string {
	var e []string
	for _, el := range a.Elements {
		e = append(e, el.String())
	}
	return fmt.Sprintf("array(%s)", strings.Join(e, ", "))
}

type ArrayElement struct {
	Key   Expression
	Value Expression
}

func (a *ArrayElement) NodeType() string { return "ArrayElement" }
func (a *ArrayElement) String() string {
	if a.Key != nil {
		return fmt.Sprintf("%s => %s", a.Key, a.Value)
	}
	return a.Value.String()
}

type EchoStatement struct {
	Values []Expression
}

func (e *EchoStatement) statementNode()   {}
func (e *EchoStatement) NodeType() string { return "EchoStatement" }
func (e *EchoStatement) String() string {
	var v []string
	for _, val := range e.Values {
		v = append(v, val.String())
	}
	return fmt.Sprintf("echo %s;", strings.Join(v, ", "))
}

type ReturnStatement struct {
	Value Expression
}

func (r *ReturnStatement) statementNode()   {}
func (r *ReturnStatement) NodeType() string { return "ReturnStatement" }
func (r *ReturnStatement) String() string {
	if r.Value != nil {
		return fmt.Sprintf("return %s;", r.Value)
	}
	return "return;"
}

type Assignment struct {
	Name  string
	Value Expression
}

func (a *Assignment) statementNode()   {}
func (a *Assignment) NodeType() string { return "Assignment" }
func (a *Assignment) String() string {
	return fmt.Sprintf("%s = %s;", a.Name, a.Value)
}

type Block struct {
	Statements []Statement
}

func (b *Block) statementNode()   {}
func (b *Block) NodeType() string { return "Block" }
func (b *Block) String() string {
	var s []string
	for _, stmt := range b.Statements {
		s = append(s, stmt.String())
	}
	return fmt.Sprintf("{\n%s}", strings.Join(s, "\n"))
}

type IfStatement struct {
	Condition Expression
	Then      *Block
	Else      []Statement
}

func (i *IfStatement) statementNode()   {}
func (i *IfStatement) NodeType() string { return "IfStatement" }
func (i *IfStatement) String() string {
	var else_ string
	if len(i.Else) > 0 {
		var s []string
		for _, stmt := range i.Else {
			s = append(s, stmt.String())
		}
		else_ = fmt.Sprintf("else {\n%s}", strings.Join(s, "\n"))
	}
	return fmt.Sprintf("if (%s) %s%s", i.Condition, i.Then, else_)
}

type ForStatement struct {
	Init      Expression
	Condition Expression
	Update    Expression
	Body      *Block
}

func (f *ForStatement) statementNode()   {}
func (f *ForStatement) NodeType() string { return "ForStatement" }
func (f *ForStatement) String() string {
	return fmt.Sprintf("for (%s; %s; %s) %s", f.Init, f.Condition, f.Update, f.Body)
}

type ForeachStatement struct {
	Expr  Expression
	Key   Expression
	Value Expression
	Body  *Block
}

func (f *ForeachStatement) statementNode()   {}
func (f *ForeachStatement) NodeType() string { return "ForeachStatement" }
func (f *ForeachStatement) String() string {
	if f.Key != nil {
		return fmt.Sprintf("foreach (%s as %s => %s) %s", f.Expr, f.Key, f.Value, f.Body)
	}
	return fmt.Sprintf("foreach (%s as %s) %s", f.Expr, f.Value, f.Body)
}

type WhileStatement struct {
	Condition Expression
	Body      *Block
}

func (w *WhileStatement) statementNode()   {}
func (w *WhileStatement) NodeType() string { return "WhileStatement" }
func (w *WhileStatement) String() string {
	return fmt.Sprintf("while (%s) %s", w.Condition, w.Body)
}

type DoWhileStatement struct {
	Body      *Block
	Condition Expression
}

func (d *DoWhileStatement) NodeType() string { return "DoWhileStatement" }
func (d *DoWhileStatement) String() string {
	return fmt.Sprintf("do %s while (%s);", d.Body, d.Condition)
}

type SwitchStatement struct {
	Condition Expression
	Cases     []SwitchCase
	Default   []Statement
}

func (s *SwitchStatement) NodeType() string { return "SwitchStatement" }
func (s *SwitchStatement) String() string {
	return fmt.Sprintf("switch (%s) { ... }", s.Condition)
}

type SwitchCase struct {
	Condition Expression
	Body      []Statement
}

func (s *SwitchCase) NodeType() string { return "SwitchCase" }
func (s *SwitchCase) String() string {
	if s.Condition != nil {
		return fmt.Sprintf("case %s:", s.Condition)
	}
	return "default:"
}

type FunctionDecl struct {
	Name       string
	Params     []FunctionParam
	ReturnType string
	Body       *Block
}

func (f *FunctionDecl) statementNode()   {}
func (f *FunctionDecl) NodeType() string { return "FunctionDecl" }
func (f *FunctionDecl) String() string {
	return fmt.Sprintf("function %s(...) { ... }", f.Name)
}

type FunctionParam struct {
	Name       string
	Type       string
	Default    Expression
	IsVariadic bool
}

func (f *FunctionParam) NodeType() string { return "FunctionParam" }
func (f *FunctionParam) String() string {
	return f.Name
}

type ClassDecl struct {
	Name       string
	Extends    string
	Implements []string
	Body       *Block
}

func (c *ClassDecl) statementNode()   {}
func (c *ClassDecl) NodeType() string { return "ClassDecl" }
func (c *ClassDecl) String() string {
	return fmt.Sprintf("class %s { ... }", c.Name)
}

type ClassMethod struct {
	Name       string
	Modifiers  []string
	Params     []FunctionParam
	ReturnType string
	Body       *Block
}

func (c *ClassMethod) NodeType() string { return "ClassMethod" }
func (c *ClassMethod) String() string {
	return fmt.Sprintf("function %s(...) { ... }", c.Name)
}

type TraitDecl struct {
	Name string
	Body *Block
}

func (t *TraitDecl) NodeType() string { return "TraitDecl" }
func (t *TraitDecl) String() string {
	return fmt.Sprintf("trait %s { ... }", t.Name)
}

type InterfaceDecl struct {
	Name    string
	Extends []string
	Methods []FunctionDecl
}

func (i *InterfaceDecl) NodeType() string { return "InterfaceDecl" }
func (i *InterfaceDecl) String() string {
	return fmt.Sprintf("interface %s { ... }", i.Name)
}

type NamespaceDecl struct {
	Name string
	Body []Statement
}

func (n *NamespaceDecl) NodeType() string { return "NamespaceDecl" }
func (n *NamespaceDecl) String() string {
	return fmt.Sprintf("namespace %s { ... }", n.Name)
}

type UseDecl struct {
	Names   []string
	Alias   string
	IsGroup bool
}

func (u *UseDecl) NodeType() string { return "UseDecl" }
func (u *UseDecl) String() string {
	return fmt.Sprintf("use %s;", strings.Join(u.Names, ", "))
}

type StaticProperty struct {
	Class  string
	Member string
}

type StaticCall struct {
	Class  string
	Method string
	Args   []Expression
}

type MethodCall struct {
	Receiver Expression
	Method   string
	Args     []Expression
}

type PropertyFetch struct {
	Receiver Expression
	Property string
}

type NewExpr struct {
	Class string
	Args  []Expression
}

func (n *NewExpr) NodeType() string { return "NewExpr" }
func (n *NewExpr) String() string {
	return fmt.Sprintf("new %s(...)", n.Class)
}

type CloneExpr struct {
	Expr Expression
}

type EmptyExpr struct {
	Expr Expression
}

type IssetExpr struct {
	Exprs []Expression
}

type EvalExpr struct {
	Code Expression
}

type ExitExpr struct {
	Code Expression
}

type GlobalStatement struct {
	Vars []Expression
}

type StaticStatement struct {
	Vars []Expression
}

type DeclareStatement struct {
	Directives map[string]Expression
	Statement  Statement
}
