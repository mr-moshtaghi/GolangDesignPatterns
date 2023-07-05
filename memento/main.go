package main

import (
	"design-patterns/memento/human"
	"design-patterns/memento/originator"
	"log"
)

func init() {
	log.SetFlags(0)
	log.Print("\n")
}

func main() {
	h := human.NewHuman()
	h.Display()

	o := originator.NewHumanOriginator(*h)
	o.Save()

	h.Damage(25)
	h.Display()
}
