package decimal

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func evalExpr(code string, env any, options []expr.Option) any {
	options = append([]expr.Option{expr.Env(env)}, options...)
	program, err := expr.Compile(code, options...)
	if err != nil {
		panic(err)
	}
	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}
	return output
}

func TestExprEnv_Decimal(t *testing.T) {
	var code string
	var env, output any

	// original result
	code = `0.1 + 0.2`
	output = evalExpr(code, env, nil)
	fmt.Println("TestExprEnv_Decimal original float:", code, "=", output)
	assert.NotEqual(t, "0.3", fmt.Sprintf("%v", output))

	// decimal result
	code = `Decimal(0.1) + Decimal(0.2)`
	env = ExprEnv{}
	output = evalExpr(code, env, GetOperatorOverloadingOptions())
	fmt.Println("TestExprEnv_Decimal decimal:", code, "=", output)
	assert.Equal(t, "0.3", fmt.Sprintf("%v", output))
}

func TestExprEnv_DecimalMul(t *testing.T) {
	var code string
	var env, output any
	code = `Decimal(9999999999999999) * Decimal(9999999999999999)`
	env = ExprEnv{}
	output = evalExpr(code, env, GetOperatorOverloadingOptions())
	fmt.Println("TestExprEnv_DecimalMul decimal:", code, "=", output)
	assert.Equal(t, "99999999999999980000000000000001", fmt.Sprintf("%v", output))
}
