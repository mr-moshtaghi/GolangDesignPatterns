package main

import (
	"design-patterns/observer/version2/notifier"
	"design-patterns/observer/version2/product"
)

func main() {
	smsSender := notifier.NewSMS()
	emailSender := notifier.NewEmail()
	ahmadClient := notifier.NewClient(1, smsSender, 5)
	sajjadClient := notifier.NewClient(2, emailSender, 6)
	product1 := product.NewProduct("shalvar")
	product2 := product.NewProduct("lebas")
	product1.AddClient(ahmadClient)
	product1.AddClient(sajjadClient)
	product1.SetNumber(6)
	product1.Broadcast()

	product2.AddClient(ahmadClient)
	product2.SetNumber(5)
	product2.Broadcast()
}
