package _6_Prototype

// Prototype 是原型接口
type Prototype interface {
	Clone() Prototype
}

// ConcretePrototype 是具体原型结构体
type ConcretePrototype struct {
	part1 string
	part2 int
}

// Clone 是克隆方法(深拷贝)
func (p ConcretePrototype) Clone() Prototype {
	return &p
}
