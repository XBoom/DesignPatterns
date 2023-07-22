package _0_Iterator

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// ConcreteIterator 具体迭代器
type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.aggregate.items)
}

func (i *ConcreteIterator) Next() interface{} {
	item := i.aggregate.items[i.index]
	i.index++
	return item
}

// Aggregate 聚合对象接口
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteAggregate 具体聚合对象
type ConcreteAggregate struct {
	items []interface{}
}

func (a *ConcreteAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{
		aggregate: a,
		index:     0,
	}
}

func (a *ConcreteAggregate) AddItem(item interface{}) {
	a.items = append(a.items, item)
}
