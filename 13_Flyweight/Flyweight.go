package _3_Flyweight

import "fmt"

// Flyweight 是享元接口
type Flyweight interface {
	Operation(extrinsicState string)
}

// ConcreteFlyweight 是具体享元类
type ConcreteFlyweight struct {
	intrinsicState string
}

func NewConcreteFlyweight(intrinsicState string) *ConcreteFlyweight {
	return &ConcreteFlyweight{
		intrinsicState: intrinsicState,
	}
}

func (f *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("具体享元对象：内部状态为 %s，外部状态为 %s\n", f.intrinsicState, extrinsicState)
}

// FlyweightFactory 是享元工厂类
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func (ff *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := ff.flyweights[key]; ok {
		return flyweight
	}

	flyweight := NewConcreteFlyweight(key)
	ff.flyweights[key] = flyweight
	return flyweight
}
