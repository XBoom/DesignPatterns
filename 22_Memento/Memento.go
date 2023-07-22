package _2_Memento

// Memento 备忘录对象
type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

// Originator 发起者对象
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.state = m.GetState()
}

// Caretaker 备忘录管理器
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}

func (c *Caretaker) SetMemento(m *Memento) {
	c.memento = m
}
