package _7_Chain

import "fmt"

// Request 请求接口
type Request interface {
	GetLevel() int
	GetContent() string
}

// ConcreteRequest 具体请求类
type ConcreteRequest struct {
	level   int
	content string
}

func NewConcreteRequest(level int, content string) *ConcreteRequest {
	return &ConcreteRequest{
		level:   level,
		content: content,
	}
}

func (r *ConcreteRequest) GetLevel() int {
	return r.level
}

func (r *ConcreteRequest) GetContent() string {
	return r.content
}

// Handler 处理者接口
type Handler interface {
	HandleRequest(request Request)
	SetNext(handler Handler)
}

// ConcreteHandler 具体处理者类
type ConcreteHandler struct {
	level int
	next  Handler
}

func NewConcreteHandler(level int) *ConcreteHandler {
	return &ConcreteHandler{
		level: level,
	}
}

func (h *ConcreteHandler) HandleRequest(request Request) {
	if request.GetLevel() == h.level {
		fmt.Printf("Handler%d 处理请求：%s\n", h.level, request.GetContent())
	} else if h.next != nil {
		h.next.HandleRequest(request)
	} else {
		fmt.Println("没有合适的处理者处理该请求")
	}
}

func (h *ConcreteHandler) SetNext(handler Handler) {
	h.next = handler
}
