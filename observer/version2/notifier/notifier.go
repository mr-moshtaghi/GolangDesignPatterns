package notifier

import "fmt"

// Observer
type MessageSender interface {
	SendMessage()
}

type Client struct {
	Id              int
	MessageSender   MessageSender
	NumberOfRequest int
}

func NewClient(id int, messageSender MessageSender, numberOfRequest int) Client {
	return Client{
		Id:              id,
		MessageSender:   messageSender,
		NumberOfRequest: numberOfRequest,
	}
}

type SMS struct{}

func NewSMS() MessageSender {
	return &SMS{}
}

func (s SMS) SendMessage() {
	fmt.Println("sending Message with SMS")
}

type Email struct{}

func NewEmail() MessageSender {
	return &Email{}
}

func (s Email) SendMessage() {
	fmt.Println("sending Message with Email")
}

type Telegram struct{}

func NewTelegram() MessageSender {
	return &Telegram{}
}

func (s Telegram) SendMessage() {
	fmt.Println("sending Message with Telegram")
}
