package main

//
//import "fmt"
//
//type Notification interface {
//	Send() string
//}
//
//type SMS struct{}
//
//func (s SMS) Send() string {
//	return "Sending SMS Notification"
//}
//
//type Email struct{}
//
//func (e Email) Send() string {
//	return "Sending Email Notification"
//}
//
//type Push struct{}
//
//func (p Push) Send() string {
//	return "Sending Push Notification"
//}
//
//func NotificationFactory(t string) Notification {
//	switch t {
//	case "sms":
//		return SMS{}
//	case "email":
//		return Email{}
//	case "push":
//		return Push{}
//	default:
//		return nil
//	}
//}
//
//func main() {
//	notification := NotificationFactory("push")
//	if notification != nil {
//		fmt.Println(notification.Send())
//	}
//}
