package _5_Observer

import "fmt"

// Observer 观察者接口
type Observer interface {
	Update(message string)
}

// Subject 被观察者接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// ConcreteObserver 具体观察者
type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("[%s] 收到消息：%s\n", co.name, message)
}

// ConcreteSubject 具体被观察者
type ConcreteSubject struct {
	observers []Observer
}

func (cs *ConcreteSubject) RegisterObserver(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) RemoveObserver(observer Observer) {
	index := -1
	for i, obs := range cs.observers {
		if obs == observer {
			index = i
			break
		}
	}
	if index >= 0 {
		cs.observers = append(cs.observers[:index], cs.observers[index+1:]...)
	}
}

func (cs *ConcreteSubject) NotifyObservers(message string) {
	for _, observer := range cs.observers {
		observer.Update(message)
	}
}
