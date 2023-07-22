package _3_Component

import (
	"fmt"
	"strings"
)

// Component 抽象构件
type Component interface {
	Add(component Component)
	Remove(component Component)
	Display(depth int)
}

// Leaf 叶子节点
type Leaf struct {
	name string
}

func (l *Leaf) Add(component Component) {}

func (l *Leaf) Remove(component Component) {}

func (l *Leaf) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + l.name)
}

// 组合节点
type Composite struct {
	name       string
	components []Component
}

func NewComposite(name string) *Composite {
	return &Composite{name, make([]Component, 0)}
}

func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

func (c *Composite) Remove(component Component) {
	for i, v := range c.components {
		if v == component {
			c.components = append(c.components[:i], c.components[i+1:]...)
		}
	}
}

func (c *Composite) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + c.name)
	for _, component := range c.components {
		component.Display(depth + 2)
	}
}
