package notification

//Concrete Product

type Email struct{}

func NewEmail() Notification {
	return &Email{}
}

func (e Email) Send() string {
	return "Sending Email Notification"
}
