package product

import (
	"design-patterns/observer/version2/notifier"
)

type Subject interface {
	AddClient(client notifier.Client)
	RemoveClient(clientId int)
	SetNumber(number int)
	Broadcast()
}

type Product struct {
	clients map[int]notifier.Client
	name    string
	number  int
}

func NewProduct(name string) Subject {
	return &Product{
		clients: make(map[int]notifier.Client),
		name:    name,
	}
}

func (p *Product) AddClient(client notifier.Client) {
	p.clients[client.Id] = client
}

func (p *Product) RemoveClient(clientId int) {
	delete(p.clients, clientId)
}

func (p *Product) SetNumber(number int) {
	p.number = number
}

func (p *Product) Broadcast() {
	for _, client := range p.clients {
		if client.NumberOfRequest == p.number {
			client.MessageSender.SendMessage()
		}
	}
}
