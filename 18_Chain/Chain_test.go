package _8_Chain

import "testing"

func TestConcreteHandler_HandleRequest(t *testing.T) {
	handler1 := NewConcreteHandler(1)
	handler2 := NewConcreteHandler(2)
	handler3 := NewConcreteHandler(3)

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	request1 := NewConcreteRequest(2, "请求 1")
	request2 := NewConcreteRequest(3, "请求 2")
	request3 := NewConcreteRequest(1, "请求 3")

	handler1.HandleRequest(request1)
	handler1.HandleRequest(request2)
	handler1.HandleRequest(request3)
}
