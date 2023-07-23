package _0_Adapter

// 目标接口
type Target interface {
	Request() string
}

// 源接口
type Adaptee interface {
	SpecificRequest() string
}

// 源接口的具体实现
type ConcreteAdaptee struct{}

func (a *ConcreteAdaptee) SpecificRequest() string {
	return "Specific request."
}

// 适配器
type Adapter struct {
	adaptee Adaptee //包含源接口
}

// 实现目标接口
func (a *Adapter) Request() string {
	return "Adapter: " + a.adaptee.SpecificRequest()
}
