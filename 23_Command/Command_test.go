package _3_Command

import "testing"

func TestInvoker_ExecuteCommand(t *testing.T) {
	receiver := &ConcreteReceiver{}
	command := NewConcreteCommand(receiver)
	invoker := NewInvoker(command)
	invoker.ExecuteCommand()
}
