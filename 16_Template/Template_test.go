package _6_Template

import "testing"

func TestConcreteClass_PrimitiveOperation2(t *testing.T) {
	GetResult(&ConcreteTemplateA{})
	GetResult(&ConcreteTemplateB{})
}
