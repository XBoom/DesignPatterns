package _1_Visitor

import "testing"

func TestObjectStructure_Accept(t *testing.T) {
	objectStructure := &ObjectStructure{}
	objectStructure.Attach(&ConcreteElementA{})
	objectStructure.Attach(&ConcreteElementB{})

	visitor := &ConcreteVisitor{}
	objectStructure.Accept(visitor)
}
