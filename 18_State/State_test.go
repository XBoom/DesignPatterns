package _8_State

import "testing"

func TestContext_Request(t *testing.T) {
	context := &Context{}

	stateA := &ConcreteStateA{}
	context.SetState(stateA)
	context.Request() // Output: 处理具体状态 A

	stateB := &ConcreteStateB{}
	context.SetState(stateB)
	context.Request() // Output: 处理具体状态 B
}
