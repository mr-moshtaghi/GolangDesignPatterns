package notification

//Concrete Product

type SMS struct{}

func NewSMS() Notification {
	return &SMS{}
}

func (s SMS) Send() string {
	return "Sending SMS Notification"
}
