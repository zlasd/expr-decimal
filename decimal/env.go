package decimal

import (
	"fmt"
	"reflect"

	"github.com/expr-lang/expr"
	decimalType "github.com/shopspring/decimal"
)

func GetOperatorOverloadingOptions() []expr.Option {
	return []expr.Option{
		expr.Operator("+", "DecimalAdd"),
		expr.Operator("-", "DecimalSub"),
		expr.Operator("*", "DecimalMul"),
		expr.Operator("/", "DecimalDiv"),
		expr.Operator("%", "DecimalMod"),
		expr.Operator("^", "DecimalPow"),
		expr.Operator("**", "DecimalPow"),
		expr.Operator("==", "DecimalEqual"),
		expr.Operator("!=", "DecimalNotEqual"),
		expr.Operator("<", "DecimalLess"),
		expr.Operator("<=", "DecimalLessEq"),
		expr.Operator(">", "DecimalGreater"),
		expr.Operator(">=", "DecimalGreaterEq"),
	}
}

type ExprEnv struct{}

func (e ExprEnv) Decimal(val any) decimalType.Decimal {
	refVal := reflect.ValueOf(val)
	switch refVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return decimalType.NewFromInt(refVal.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return decimalType.NewFromInt(int64(refVal.Uint()))
	case reflect.Float32, reflect.Float64:
		return decimalType.NewFromFloat(refVal.Float())
	case reflect.String:
		return decimalType.RequireFromString(refVal.String())
	default:
		panic(fmt.Sprintf("invalid type: %T", val))
	}
}

func (e ExprEnv) DecimalAdd(a, b decimalType.Decimal) decimalType.Decimal {
	return a.Add(b)
}

func (e ExprEnv) DecimalSub(a, b decimalType.Decimal) decimalType.Decimal {
	return a.Sub(b)
}

func (e ExprEnv) DecimalMul(a, b decimalType.Decimal) decimalType.Decimal {
	return a.Mul(b)
}

func (e ExprEnv) DecimalDiv(a, b decimalType.Decimal) decimalType.Decimal {
	return a.Div(b)
}

func (e ExprEnv) DecimalMod(a, b decimalType.Decimal) decimalType.Decimal {
	return a.Mod(b)
}

func (e ExprEnv) DecimalPow(a, b decimalType.Decimal) decimalType.Decimal {
	return a.Pow(b)
}

func (e ExprEnv) DecimalEqual(a, b decimalType.Decimal) bool {
	return a.Equal(b)
}

func (e ExprEnv) DecimalNotEqual(a, b decimalType.Decimal) bool {
	return !a.Equal(b)
}

func (e ExprEnv) DecimalLess(a, b decimalType.Decimal) bool {
	return a.LessThan(b)
}

func (e ExprEnv) DecimalLessEq(a, b decimalType.Decimal) bool {
	return a.LessThanOrEqual(b)
}

func (e ExprEnv) DecimalGreater(a, b decimalType.Decimal) bool {
	return a.GreaterThan(b)
}

func (e ExprEnv) DecimalGreaterEq(a, b decimalType.Decimal) bool {
	return a.GreaterThanOrEqual(b)
}
