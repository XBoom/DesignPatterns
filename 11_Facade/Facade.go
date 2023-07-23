package _1_Facade

import "fmt"

// ComponentA 子系统中的组件
type ComponentA struct{}

func (c *ComponentA) MethodA() {
	fmt.Println("ComponentA: MethodA")
}

type ComponentB struct{}

func (c *ComponentB) MethodB() {
	fmt.Println("ComponentB: MethodB")
}

type ComponentC struct{}

func (c *ComponentC) MethodC() {
	fmt.Println("ComponentC: MethodC")
}

// Facade 门面
type Facade struct {
	componentA *ComponentA
	componentB *ComponentB
	componentC *ComponentC
}

func NewFacade() *Facade {
	return &Facade{
		componentA: &ComponentA{},
		componentB: &ComponentB{},
		componentC: &ComponentC{},
	}
}

// Operation 提供简单统一的接口给客户端使用
func (f *Facade) Operation() {
	fmt.Println("Facade: Operation")
	f.componentA.MethodA()
	f.componentB.MethodB()
	f.componentC.MethodC()
}
