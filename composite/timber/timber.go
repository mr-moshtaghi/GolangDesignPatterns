package timber

import (
	timber_interface "design-patterns/composite/interface"
	log "github.com/sirupsen/logrus"
)

func Print(tree timber_interface.NodeTree) {
	doPrint(tree, " ")
}

func doPrint(tree timber_interface.NodeTree, prefix string) {
	log.Print(prefix, tree.Display())

	for _, child := range tree.Components() {
		doPrint(child, prefix+" ")
	}
}
