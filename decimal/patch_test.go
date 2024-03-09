package decimal

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExprPatch(t *testing.T) {
	var code string
	var output any

	code = `0.1 + 0.2`
	options := GetOperatorOverloadingOptions()
	options = append([]expr.Option{expr.Patch(&ExprPatch{})}, options...)
	output = evalExpr(code, ExprEnv{}, options)
	fmt.Println("TestExprPatch", code, "=", output)
	assert.Equal(t, "0.3", fmt.Sprintf("%v", output))
}
