package _5_Mediator

import "testing"

func TestConcreteMediator_SendMessage(t *testing.T) {
	//1. 中介对象包含具体对象
	mediator := &ConcreteMediator{}
	colleague1 := &ConcreteColleague{name: "Colleague1", mediator: mediator}
	colleague2 := &ConcreteColleague{name: "Colleague2", mediator: mediator}
	//2. 中介双方
	mediator.colleague1 = colleague1
	mediator.colleague2 = colleague2
	//3. 发送消息
	colleague1.SendMessage("Hello, colleague2!")
	colleague2.SendMessage("Hi, colleague1!")
}
