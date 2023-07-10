package _8_Bridge

import (
	"testing"
)

func TestConcreteImplementorA_operationImpl(t *testing.T) {
	implementorA := &ConcreteImplementorA{}
	abstractionA := NewRefinedAbstraction(implementorA)
	resultA := abstractionA.Operation()
	if resultA != "ConcreteImplementorA" {
		t.Fatalf("except is not %s", resultA)
	}
}

func TestConcreteImplementorB_operationImpl(t *testing.T) {
	implementorB := &ConcreteImplementorB{}
	abstractionB := NewRefinedAbstraction(implementorB)
	resultA := abstractionB.Operation()
	if resultA != "ConcreteImplementorB" {
		t.Fatalf("except is not %s", resultA)
	}
}
