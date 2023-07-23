package _2_Command

import "fmt"

// Command 命令接口
type Command interface {
	Execute()
}

// ConcreteCommand 具体命令
type ConcreteCommand struct {
	receiver Receiver
}

func NewConcreteCommand(receiver Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver: receiver,
	}
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

// Receiver 接收者接口
type Receiver interface {
	Action()
}

// ConcreteReceiver 具体接收者
type ConcreteReceiver struct{}

func (cr *ConcreteReceiver) Action() {
	fmt.Println("接收者执行具体操作")
}

// Invoker 调用者
type Invoker struct {
	command Command
}

func NewInvoker(command Command) *Invoker {
	return &Invoker{
		command: command,
	}
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
