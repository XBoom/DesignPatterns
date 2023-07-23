package _9_Iterator

import (
	"fmt"
	"testing"
)

func TestConcreteIterator_Next(t *testing.T) {
	aggregate := &ConcreteAggregate{}
	aggregate.AddItem("Item 1")
	aggregate.AddItem("Item 2")
	aggregate.AddItem("Item 3")

	iterator := aggregate.CreateIterator()
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}
