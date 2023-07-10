package _8_Bridge

// Implementor 具体实现类接口
type Implementor interface {
	operationImpl() string
}

type ConcreteImplementorA struct{}

func (c *ConcreteImplementorA) operationImpl() string {
	return "ConcreteImplementorA"
}

var _ Implementor = (*ConcreteImplementorA)(nil)

type ConcreteImplementorB struct{}

func (c *ConcreteImplementorB) operationImpl() string {
	return "ConcreteImplementorB"
}

var _ Implementor = (*ConcreteImplementorB)(nil)

// Abstraction 抽象类
type Abstraction struct {
	implementor Implementor
}

func (a *Abstraction) Operation() string {
	return a.implementor.operationImpl()
}

// RefinedAbstraction 抽象类的扩展类
type RefinedAbstraction struct {
	Abstraction
}

func NewRefinedAbstraction(implementor Implementor) *RefinedAbstraction {
	return &RefinedAbstraction{Abstraction{implementor}}
}
