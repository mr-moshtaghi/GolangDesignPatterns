package notification

//Concrete Product

type Push struct{}

func NewPush() Notification {
	return &Push{}
}

func (p Push) Send() string {
	return "Sending Push Notification"
}
