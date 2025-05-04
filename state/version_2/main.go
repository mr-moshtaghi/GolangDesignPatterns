package main

import "log"

type State interface {
	think() string
}

type Happy struct{}

func (h *Happy) think() string {
	return "Everything is great! I'm feeling positive and optimistic."
}

type Sad struct{}

func (s Sad) think() string {
	return "I'm feeling down. Everything seems gloomy and hard."
}

type Angry struct{}

func (a Angry) think() string {
	return "I'm frustrated and angry. Everything is annoying me right now."
}

type Human struct {
	state State
}

func (h *Human) setState(state State) {
	h.state = state
}

func (h *Human) think() string {
	return h.state.think()
}
func main() {
	person := Human{
		state: &Happy{},
	}
	log.Println(person.think())

	person.setState(Sad{})
	log.Println(person.think())

	person.setState(Angry{})
	log.Println(person.think())
}
