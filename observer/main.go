package main

import (
	"fmt"
	"log"
	"math/rand"
)

// interface
type Publisher interface {
	addSubscriber(subscriber Subscriber)
	removeSubscriber(subId string)
	broadcast(msg string)
}

type Subscriber interface {
	id() string
	react(msg string)
}

// implementation > Publisher
type publisher struct {
	subscribers map[string]Subscriber
}

func newPublisher() publisher {
	return publisher{subscribers: make(map[string]Subscriber)}
}

func (p publisher) addSubscriber(subscriber Subscriber) {
	p.subscribers[subscriber.id()] = subscriber
}

func (p publisher) removeSubscriber(subId string) {
	delete(p.subscribers, subId)
}

func (p publisher) broadcast(msg string) {
	for _, subscriber := range p.subscribers {
		subscriber.react(msg)
	}
}

// implementation > Subscriber
type subscriber struct {
	subId string
}

func newSubscriber(subId string) subscriber {
	return subscriber{subId: subId}
}

func (s subscriber) id() string {
	return s.subId
}

func (s subscriber) react(msg string) {
	log.Printf("ID %s - received: %s", s.subId, msg)
}

// implementation > Auto-generated ID Subscriber
type autogenerateIdSubscriber struct {
	subId string
}

func newAutogenerateIdSubscriber() *autogenerateIdSubscriber {
	return &autogenerateIdSubscriber{subId: fmt.Sprint(rand.Int())}
}

func (a autogenerateIdSubscriber) id() string {
	return a.subId
}

func (a autogenerateIdSubscriber) react(msg string) {
	log.Printf("ID %s - auto generatedId sub received: %s", a.subId, msg)
}

func main() {
	var p Publisher
	p = newPublisher()
	p.broadcast("hello")

	s := newSubscriber("321")
	s2 := newSubscriber("456")
	p.addSubscriber(s)
	p.addSubscriber(s2)
	p.broadcast("hello again")

	p.removeSubscriber(s.id())
	p.broadcast("good afternoon")

	a := newAutogenerateIdSubscriber()
	p.addSubscriber(a)
	p.broadcast("bye")
}
