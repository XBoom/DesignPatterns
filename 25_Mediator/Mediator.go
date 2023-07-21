package _5_Mediator

import "fmt"

// Colleague 同伴类接口
type Colleague interface {
	SendMessage(message string)
	ReceiveMessage(message string)
}

// Mediator 中介者接口
type Mediator interface {
	SendMessage(message string, colleague Colleague)
}

// ConcreteMediator 具体中介者
type ConcreteMediator struct {
	colleague1 Colleague
	colleague2 Colleague
}

func (c *ConcreteMediator) SendMessage(message string, colleague Colleague) {
	if colleague == c.colleague1 {
		c.colleague2.ReceiveMessage(message)
	} else {
		c.colleague1.ReceiveMessage(message)
	}
}

// ConcreteColleague 具体同伴类
type ConcreteColleague struct {
	name     string
	mediator Mediator
}

// SendMessage 发送消息
func (c *ConcreteColleague) SendMessage(message string) {
	c.mediator.SendMessage(message, c)
}

// ReceiveMessage 接受消息
func (c *ConcreteColleague) ReceiveMessage(message string) {
	fmt.Printf("%s 收到消息：%s\n", c.name, message)
}
