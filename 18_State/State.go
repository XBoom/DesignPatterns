package _8_State

import "fmt"

// State 状态接口
type State interface {
	HandleState()
}

// ConcreteStateA 具体状态 A
type ConcreteStateA struct{}

func (s *ConcreteStateA) HandleState() {
	fmt.Println("处理具体状态 A")
}

// ConcreteStateB 具体状态 B
type ConcreteStateB struct{}

func (s *ConcreteStateB) HandleState() {
	fmt.Println("处理具体状态 B")
}

// Context 环境类
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.HandleState()
}
