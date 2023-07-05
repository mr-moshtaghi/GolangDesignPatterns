package memento

import "design-patterns/memento/human"

// interface
type Memento interface {
	Restore() interface{}
}

// human implementation
type HumanMemento struct {
	human human.Human
}

func NewHumanMemento(human human.Human) *HumanMemento {
	return &HumanMemento{human: human}
}

func (memento HumanMemento) Restore() interface{} {
	return memento.human
}
