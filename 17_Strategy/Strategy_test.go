package _7_Strategy

import (
	"fmt"
	"testing"
)

func TestContext_ExecuteStrategy(t *testing.T) {
	context := &Context{}

	// 使用加法策略
	context.SetStrategy(&ConcreteStrategyAdd{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Println("加法策略结果:", result) // Output: 加法策略结果: 15

	// 使用减法策略
	context.SetStrategy(&ConcreteStrategySubtract{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Println("减法策略结果:", result) // Output: 减法策略结果: 5
}
