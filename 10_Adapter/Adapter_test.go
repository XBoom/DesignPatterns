package _0_Adapter

import (
	"fmt"
	"testing"
)

func TestAdapter_Request(t *testing.T) {
	adaptee := &ConcreteAdaptee{}
	adapter := &Adapter{adaptee}

	result := adapter.Request()
	fmt.Println(result)
}
