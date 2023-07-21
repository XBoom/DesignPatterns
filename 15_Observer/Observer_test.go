package _5_Observer

import "testing"

func TestConcreteSubject_NotifyObservers(t *testing.T) {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}

	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	subject.NotifyObservers("Hello, observers!")

	subject.RemoveObserver(observer1)

	subject.NotifyObservers("Hi, observers!")
}
