package _6_Prototype

import (
	"fmt"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	prototype := &ConcretePrototype{
		part1: "hello",
		part2: 0,
	}
	clone1 := prototype.Clone()
	prototype.part1 = "world"
	prototype.part2 = 100
	clone2 := prototype.Clone()
	fmt.Printf("p1 %p %+v \n", &prototype, prototype)
	fmt.Printf("p2 %p %+v \n", &clone1, clone1)
	fmt.Printf("p3 %p %+v \n", &clone2, clone2)
}
