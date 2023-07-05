package originator

import (
	"design-patterns/memento/human"
	"design-patterns/memento/memento"
)

type HumanOriginator struct {
	human human.Human
}

func NewHumanOriginator(human human.Human) *HumanOriginator {
	return &HumanOriginator{human: human}
}

func (originator *HumanOriginator) Save() memento.Memento {
	return memento.NewHumanMemento(originator.human)
}

func (originator *HumanOriginator) Restore(m memento.Memento) {
	originator.human = m.Restore().(human.Human)
}
