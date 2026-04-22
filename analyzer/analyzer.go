package analyzer

import (
	"fmt"

	"github.com/phat9k/internal/ast"
)

type Symbol struct {
	Name string
	Kind SymbolKind
	Type string
}

type SymbolKind int

const (
	KindVariable SymbolKind = iota
	KindFunction
	KindClass
	KindMethod
	KindProperty
	KindConstant
)

type Error struct {
	Message  string
	Line     int
	Column   int
	Severity Severity
}

type Severity int

const (
	Warning Severity = iota
	Error_
	Fatal
)

type Metrics struct {
	Functions    int
	Classes      int
	Methods     int
	Statements  int
	Complexity  int
}

type Analyzer struct {
	program *ast.Program
	Symbols map[string]*Symbol
	Errors  []Error
	Metrics Metrics
}

func New(program *ast.Program) *Analyzer {
	return &Analyzer{
		program: program,
		Symbols: make(map[string]*Symbol),
	}
}

func (a *Analyzer) Analyze() (*AnalysisResult, error) {
	a.collectDeclarations()
	a.calculateMetrics()

	return &AnalysisResult{
		Symbols: a.Symbols,
		Errors:  a.Errors,
		Metrics: a.Metrics,
	}, nil
}

type AnalysisResult struct {
	Symbols map[string]*Symbol
	Errors  []Error
	Metrics Metrics
}

func (a *Analyzer) collectDeclarations() {
	for _, stmt := range a.program.Body {
		a.collect(stmt)
	}
}

func (a *Analyzer) collect(node ast.Node) {
	switch n := node.(type) {
	case *ast.FunctionDecl:
		a.Symbols[n.Name] = &Symbol{Name: n.Name, Kind: KindFunction}
		a.Metrics.Functions++

	case *ast.ClassDecl:
		a.Symbols[n.Name] = &Symbol{Name: n.Name, Kind: KindClass}
		a.Metrics.Classes++

	case *ast.Assignment:
		a.Symbols[n.Name] = &Symbol{Name: n.Name, Kind: KindVariable}

	case *ast.Block:
		for _, s := range n.Statements {
			a.collect(s)
		}
	}
}

func (a *Analyzer) calculateMetrics() {
	for _, stmt := range a.program.Body {
		a.count(stmt)
	}
}

func (a *Analyzer) count(node ast.Node) {
	a.Metrics.Statements++
	switch n := node.(type) {
	case *ast.Block:
		for _, s := range n.Statements {
			a.count(s)
		}
	case *ast.IfStatement:
		a.Metrics.Complexity++
	case *ast.ForStatement:
		a.Metrics.Complexity++
	case *ast.WhileStatement:
		a.Metrics.Complexity++
	case *ast.ForeachStatement:
		a.Metrics.Complexity++
	case *ast.SwitchStatement:
		a.Metrics.Complexity++
	}
}

func FormatResult(r *AnalysisResult) string {
	out := fmt.Sprintf("Functions: %d\n", r.Metrics.Functions)
	out += fmt.Sprintf("Classes: %d\n", r.Metrics.Classes)
	out += fmt.Sprintf("Statements: %d\n", r.Metrics.Statements)
	out += fmt.Sprintf("Complexity: %d\n", r.Metrics.Complexity)
	out += "\nSymbols:\n"
	for name, sym := range r.Symbols {
		out += fmt.Sprintf("  %s: %v\n", name, sym.Kind)
	}
	return out
}