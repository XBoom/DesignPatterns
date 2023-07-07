package _0_Decorator

// Component 抽象构件，给这个对象动态地添加职责
type Component interface {
	Operation() string
}

// ConcreteComponent 具体构件
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator 装饰者
type Decorator interface {
	Component
}

// ConcreteDecoratorA 具体装饰者 A
type ConcreteDecoratorA struct {
	component Component
}

func (d *ConcreteDecoratorA) Operation() string {
	return "ConcreteDecoratorA(" + d.component.Operation() + ")"
}

// ConcreteDecoratorB 具体装饰者 B
type ConcreteDecoratorB struct {
	component Component
}

func (d *ConcreteDecoratorB) Operation() string {
	return "ConcreteDecoratorB(" + d.component.Operation() + ")"
}
