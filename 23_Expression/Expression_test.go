package _3_Expression

import (
	"fmt"
	"testing"
)

func TestNonTerminalExpression_Interpret(t *testing.T) {
	// 构建解释器规则：(true and false) or (true or true)
	exp1 := NewTerminalExpression("true")
	exp2 := NewTerminalExpression("false")
	andExpr := NewNonTerminalExpression(exp1, exp2, "and")

	exp3 := NewTerminalExpression("true")
	exp4 := NewTerminalExpression("true")
	orExpr := NewNonTerminalExpression(exp3, exp4, "or")

	rootExpr := NewNonTerminalExpression(andExpr, orExpr, "or")

	// 解释并计算表达式结果
	result := rootExpr.Interpret()
	fmt.Println("表达式计算结果：", result) // Output: 表达式计算结果： true
}
