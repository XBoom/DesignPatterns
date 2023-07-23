package _1_Memento

import (
	"fmt"
	"testing"
)

func TestCaretaker_SetMemento(t *testing.T) {
	//1. 原始对象设置状态
	originator := &Originator{}
	originator.SetState("State 1")
	fmt.Println("当前状态:", originator.state)

	//2. 备忘发起者
	caretaker := &Caretaker{}
	// 创建备忘录并保存状态
	caretaker.SetMemento(originator.CreateMemento())

	originator.SetState("State 2")
	fmt.Println("更新后状态:", originator.state)

	//3. 恢复状态
	originator.RestoreMemento(caretaker.GetMemento())
	fmt.Println("恢复后状态:", originator.state)
}
