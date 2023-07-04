package main

import "log"

type human struct {
	age      int
	height   int
	eyeColor string
}

// Empty constructor
func newHuman() human {
	return human{}
}

func newHumanWithFields(age int, height int, eyeColor string) human {
	return human{
		age:      age,
		height:   height,
		eyeColor: eyeColor,
	}
}

// Builder
func (h human) withEyeColor(color string) human {
	h.eyeColor = color

	return h
}

func (h human) withAge(age int) human {
	h.age = age

	return h
}

func (h human) withHeight(height int) human {
	h.height = height

	return h
}

// reset
func (h human) reset() human {
	h = newHuman()
	return h
}

// Pre-defined builders
func giant() human {
	return newHuman().withHeight(280).withEyeColor("green")
}

func main() {
	me := newHuman().withEyeColor("black").withAge(25).withHeight(180).reset()
	you := newHumanWithFields(26, 170, "blue")

	log.Printf("%+v", me)
	log.Printf("%+v", you)

	youngGiant := giant().withAge(20)
	oldGiant := giant().withAge(90)

	log.Printf("%+v", youngGiant)
	log.Printf("%+v", oldGiant)

}
