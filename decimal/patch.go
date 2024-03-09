package decimal

import (
	"github.com/expr-lang/expr/ast"
	"reflect"
)

type ExprPatch struct {
	patched map[uintptr]bool

	DecimalVarMap map[string]bool
}

func (p *ExprPatch) DecimalVar(vars []string) {
	p.ensureInitialized()
	for _, v := range vars {
		p.DecimalVarMap[v] = true
	}
}

func (p *ExprPatch) Visit(node *ast.Node) {
	p.ensureInitialized()
	refVal := reflect.ValueOf(*node)
	if p.patched[refVal.Pointer()] {
		return
	}
	p.patched[refVal.Pointer()] = true
	switch callNode := (*node).(type) {
	case *ast.IntegerNode, *ast.FloatNode:
		ast.Patch(node, &ast.CallNode{
			Callee:    &ast.IdentifierNode{Value: "Decimal"},
			Arguments: []ast.Node{*node},
		})
	case *ast.IdentifierNode:
		if p.DecimalVarMap[callNode.Value] {
			ast.Patch(node, &ast.CallNode{
				Callee:    &ast.IdentifierNode{Value: "Decimal"},
				Arguments: []ast.Node{*node},
			})
		}
	}
}

func (p *ExprPatch) ensureInitialized() {
	if p.patched == nil {
		p.patched = make(map[uintptr]bool)
	}
	if p.DecimalVarMap == nil {
		p.DecimalVarMap = make(map[string]bool)
	}
}
