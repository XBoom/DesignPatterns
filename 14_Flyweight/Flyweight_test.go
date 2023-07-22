package _4_Flyweight

import "testing"

func TestConcreteFlyweight_Operation(t *testing.T) {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("A")
	flyweight1.Operation("1")

	flyweight2 := factory.GetFlyweight("B")
	flyweight2.Operation("2")

	flyweight3 := factory.GetFlyweight("A")
	flyweight3.Operation("3")
}
