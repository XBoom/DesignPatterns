package _7_Strategy

// Strategy 策略接口
type Strategy interface {
	Execute(num1, num2 int) int
}

// ConcreteStrategyAdd 具体策略类 - 加法
type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) Execute(num1, num2 int) int {
	return num1 + num2
}

// ConcreteStrategySubtract 具体策略类 - 减法
type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) Execute(num1, num2 int) int {
	return num1 - num2
}

// Context 上下文类
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.Execute(num1, num2)
}
