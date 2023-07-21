package _5_Mediator

// Colleague 同事类接口
type Colleague interface {
	SendMessage(message string)
	ReceiveMessage(message string)
}
