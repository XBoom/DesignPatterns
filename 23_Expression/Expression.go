package _3_Expression

// Expression 表达式接口
type Expression interface {
	Interpret() bool
}

// TerminalExpression 终结符表达式
type TerminalExpression struct {
	data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{
		data: data,
	}
}

func (t *TerminalExpression) Interpret() bool {
	// 在实际应用中，此处可以进行语言文法解释的具体逻辑判断
	// 这里简化，当 data 为 "true" 时返回 true，否则返回 false
	return t.data == "true"
}

// NonTerminalExpression 非终结符表达式
type NonTerminalExpression struct {
	exp1     Expression
	exp2     Expression
	operator string
}

func NewNonTerminalExpression(exp1, exp2 Expression, operator string) *NonTerminalExpression {
	return &NonTerminalExpression{
		exp1:     exp1,
		exp2:     exp2,
		operator: operator,
	}
}

func (n *NonTerminalExpression) Interpret() bool {
	// 在实际应用中，此处可以进行语言文法解释的具体逻辑判断
	// 这里简化，只实现简单的逻辑运算，支持 "and" 和 "or" 操作符
	switch n.operator {
	case "and":
		return n.exp1.Interpret() && n.exp2.Interpret()
	case "or":
		return n.exp1.Interpret() || n.exp2.Interpret()
	default:
		return false
	}
}
