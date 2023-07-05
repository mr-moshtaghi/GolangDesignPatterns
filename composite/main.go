package main

import (
	timber_interface "design-patterns/composite/interface"
	"design-patterns/composite/timber"
)

// for example
type Leaf struct {
	label string
}

func (l Leaf) Display() string {
	return l.label
}

func (l Leaf) Components() []timber_interface.NodeTree {
	return nil
}

type Branch struct {
	label      string
	components []timber_interface.NodeTree
}

func (b Branch) Display() string {
	return b.label
}

func (b Branch) Components() []timber_interface.NodeTree {
	return b.components
}

func main() {
	l0 := Leaf{"L0"}
	l1 := Leaf{"L1"}

	b1 := Branch{
		label: "branch1",
		components: []timber_interface.NodeTree{
			Leaf{"l3"},
		},
	}

	b0 := Branch{
		label: "branch0",
		components: []timber_interface.NodeTree{
			l0, l1, b1,
		},
	}

	timber.Print(b0)
}
