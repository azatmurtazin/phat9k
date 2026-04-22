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
	Message string
	Line   int
	Column int
}

type Metrics struct {
	Functions   int
	Classes     int
	Methods    int
	LinesOfCode int
	Statements int
}

type Analyzer struct {
	program *ast.Program
	Symbols map[string]*Symbol
	Errors  []Error
	Metrics Metrics
	types   map[string]string
}

func New(program *ast.Program) *Analyzer {
	return &Analyzer{
		program: program,
		Symbols: make(map[string]*Symbol),
		types:   make(map[string]string),
	}
}

func (a *Analyzer) Analyze() (*AnalysisResult, error) {
	a.collectDeclarations()
	a.calculateMetrics()

	return &AnalysisResult{
		Symbols: a.Symbols,
		Errors:  a.Errors,
		Metrics: a.Metrics,
		Types:   a.types,
	}, nil
}

type AnalysisResult struct {
	Symbols map[string]*Symbol
	Errors  []Error
	Metrics Metrics
	Types   map[string]string
}

func (a *Analyzer) collectDeclarations() {
	for _, stmt := range a.program.Body {
		a.collectDecl(stmt)
	}
}

func (a *Analyzer) collectDecl(node ast.Statement) {
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
		for _, stmt := range n.Statements {
			a.collectDecl(stmt)
		}
	}
}

func (a *Analyzer) calculateMetrics() {
	for _, stmt := range a.program.Body {
		a.countStatement(stmt)
	}
}

func (a *Analyzer) countStatement(node ast.Statement) {
	a.Metrics.Statements++
	switch n := node.(type) {
	case *ast.Block:
		for _, stmt := range n.Statements {
			a.countStatement(stmt)
		}
	case *ast.IfStatement:
		if n.Then != nil {
			a.countStatement(n.Then)
		}
	case *ast.ForStatement:
		if n.Body != nil {
			a.countStatement(n.Body)
		}
	case *ast.WhileStatement:
		if n.Body != nil {
			a.countStatement(n.Body)
		}
	case *ast.ForeachStatement:
		if n.Body != nil {
			a.countStatement(n.Body)
		}
	case *ast.FunctionDecl:
		if n.Body != nil {
			a.countStatement(n.Body)
		}
	case *ast.ClassDecl:
		if n.Body != nil {
			a.countStatement(n.Body)
		}
	}
}

func (a *Analyzer) addError(msg string, line, col int) {
	a.Errors = append(a.Errors, Error{
		Message: msg,
		Line:    line,
		Column: col,
	})
}

func FormatResult(result *AnalysisResult) string {
	output := fmt.Sprintf("Symbols: %d\n", len(result.Symbols))
	output += fmt.Sprintf("Errors: %d\n", len(result.Errors))
	output += fmt.Sprintf("Functions: %d\n", result.Metrics.Functions)
	output += fmt.Sprintf("Classes: %d\n", result.Metrics.Classes)
	output += fmt.Sprintf("Statements: %d\n", result.Metrics.Statements)
	for name, sym := range result.Symbols {
		output += fmt.Sprintf("  %s: %v\n", name, sym.Kind)
	}
	for _, err := range result.Errors {
		output += fmt.Sprintf("Error at line %d: %s\n", err.Line, err.Message)
	}
	return output
}