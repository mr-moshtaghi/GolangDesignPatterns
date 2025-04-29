package main

import (
	notification "design-patterns/factory_method_di/notification"
	"fmt"
)

func main() {
	factory := notification.NewNotificationFactory()
	factory.Register("SMS", notification.NewSMS)
	factory.Register("Email", notification.NewEmail)
	factory.Register("Push", notification.NewPush)

	sms, err := factory.Create("SMS")
	fmt.Println(err)
	fmt.Println(sms.Send())
}
