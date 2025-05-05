package main

import (
	"errors"
	"fmt"
	"time"
)

type InventoryService struct{}

func (is *InventoryService) CheckStock(productID string, quantity int) error {
	fmt.Printf("Inventory: Checking stock for Product %s, Qty %d...\n", productID, quantity)
	if quantity > 10 {
		return errors.New("inventory: not enough stock available")
	}
	fmt.Println("Inventory: Stock available.")
	return nil
}

func (is *InventoryService) DecreaseStock(productID string, quantity int) error {
	fmt.Printf("Inventory: Decreasing stock for Product %s, Qty %d...\n", productID, quantity)
	fmt.Println("Inventory: Stock decreased.")
	return nil
}

func (is *InventoryService) IncreaseStock(productID string, quantity int) error {
	fmt.Printf("Inventory: Increasing stock for Product %s, Qty %d...\n", productID, quantity)
	fmt.Println("Inventory: Stock Increased.")
	return nil
}

type PaymentGateway struct{}

func (pg *PaymentGateway) ChargeCard(cardNumber string, amount float64) error {
	fmt.Printf("Payment: Charging %.2f using card %s...\n", amount, cardNumber)
	if amount > 5000.0 {
		return errors.New("payment: transaction amount exceeds limit")
	}
	fmt.Println("Payment: Payment successful.")
	return nil
}

func (pg *PaymentGateway) RefundCharge(cardNumber string, amount float64) error {
	fmt.Printf("Payment: Refunding %.2f to card %s...\n", amount, cardNumber)
	fmt.Println("Payment: Refund successful.")
	return nil
}

type ShippingService struct{}

func (s *ShippingService) ScheduleDelivery(productID string, quantity int, address string) error {
	fmt.Printf("Shipping: Scheduling delivery for %d of %s to %s...\n", quantity, productID, address)
	fmt.Println("Shipping: Delivery scheduled.")
	return nil
}

type NotificationService struct{}

func (ns *NotificationService) SendOrderConfirmation(customerEmail string, orderID string) error {
	fmt.Printf("Notification: Sending order confirmation email to %s for order %s...\n", customerEmail, orderID)
	fmt.Println("Notification: Email sent.")
	return nil
}

type OrderRepository struct{}

func (or *OrderRepository) SaveOrder(orderID string, productID string, quantity int, customerEmail string) error {
	fmt.Printf("Repository: Saving order %s for %s (Qty %d) for customer %s...\n", orderID, productID, quantity, customerEmail)
	fmt.Println("Repository: Order saved.")
	return nil
}

type CustomerInfo struct {
	Email   string
	Address string
}

type PaymentDetails struct {
	CardNumber string
	Amount     float64
}

type OrderFacade struct {
	inventory    *InventoryService
	payment      *PaymentGateway
	shipping     *ShippingService
	notification *NotificationService
	repository   *OrderRepository
}

func NewOrderFace() *OrderFacade {
	return &OrderFacade{
		inventory:    &InventoryService{},
		payment:      &PaymentGateway{},
		shipping:     &ShippingService{},
		notification: &NotificationService{},
		repository:   &OrderRepository{},
	}
}

func (of *OrderFacade) PlaceOrder(
	productID string, quantity int, customer CustomerInfo, payment PaymentDetails,
) (string, error) {
	fmt.Println("\n--- Order Facade: Starting PlaceOrder Process ---")

	err := of.inventory.CheckStock(productID, quantity)
	if err != nil {
		fmt.Println("Order Facade: Inventory check failed.")
		return "", fmt.Errorf("order placement failed: inventory check: %w", err)
	}

	err = of.payment.ChargeCard(payment.CardNumber, payment.Amount)
	if err != nil {
		fmt.Println("Order Facade: Payment processing failed.")
		return "", fmt.Errorf("order placement failed: payment process: %w", err)
	}
	err = of.inventory.DecreaseStock(productID, quantity)
	if err != nil {
		fmt.Println("Order Facade: Decreasing stock failed.")
		_ = of.payment.RefundCharge(payment.CardNumber, payment.Amount)
		return "", fmt.Errorf("order placement failed: decrease stock: %w", err)
	}

	err = of.shipping.ScheduleDelivery(productID, quantity, customer.Address)
	if err != nil {
		fmt.Println("Order Facade: Shipping scheduling failed.")
		_ = of.inventory.IncreaseStock(productID, quantity)
		_ = of.payment.RefundCharge(payment.CardNumber, payment.Amount)
		return "", fmt.Errorf("order placement failed: shipping schedule: %w", err)
	}

	orderID := fmt.Sprintf("ORDER-%d", time.Now().UnixNano())
	err = of.repository.SaveOrder(orderID, productID, quantity, customer.Email)
	if err != nil {
		fmt.Println("Order Facade: Saving order failed.")
		_ = of.inventory.IncreaseStock(productID, quantity)
		_ = of.payment.RefundCharge(payment.CardNumber, payment.Amount)
		return "", fmt.Errorf("order placement failed: save order: %w", err)
	}

	err = of.notification.SendOrderConfirmation(customer.Email, orderID)
	if err != nil {
		fmt.Println("Order Facade: Sending notification failed (order might still be valid).")
	}
	fmt.Printf("--- Order Facade: Order %s Placed Successfully ---\n", orderID)
	return orderID, nil
}

func main() {
	orderFacade := NewOrderFace()
	OrderSuccessful(orderFacade)
	fmt.Println("--- Client: Attempting a successful order ---")

	OrderFailLowStock(orderFacade)
	fmt.Println("\n--- Client: Attempting an order that fails due to low stock ---")

	OrderFailHighPaymentAmount(orderFacade)
	fmt.Println("\n--- Client: Attempting an order that fails due to high payment amount ---")

}

func OrderSuccessful(orderFacade *OrderFacade) {
	successfulOrderID, err := orderFacade.PlaceOrder(
		"LAPTOP-XYZ",
		1,
		CustomerInfo{Email: "alice@example.com", Address: "101 High Street, Cityville"},
		PaymentDetails{CardNumber: "1111-2222-3333-4444", Amount: 1200.0},
	)
	if err != nil {
		fmt.Printf("Client received error for successful order attempt: %v\n", err)
	} else {
		fmt.Printf("Client successfully placed order with ID: %s\n", successfulOrderID)
	}
}

func OrderFailLowStock(orderFacade *OrderFacade) {
	_, err := orderFacade.PlaceOrder(
		"GADGET-PRO",
		15,
		CustomerInfo{Email: "bob@example.com", Address: "202 Low Road, Villagetown"},
		PaymentDetails{CardNumber: "5555-6666-7777-8888", Amount: 500.0},
	)
	if err != nil {
		fmt.Printf("Client received expected error for low stock order: %v\n", err)
	} else {
		fmt.Println("Client unexpectedly placed order (should have failed).")
	}

}

func OrderFailHighPaymentAmount(orderFacade *OrderFacade) {
	_, err := orderFacade.PlaceOrder(
		"SERVER-RACK",
		1,
		CustomerInfo{Email: "charlie@example.com", Address: "303 Big Building, Metropolis"},
		PaymentDetails{CardNumber: "9999-8888-7777-6666", Amount: 7000.0}, // مبلغ بالا
	)

	if err != nil {
		fmt.Printf("Client received expected error for high payment amount: %v\n", err)
	} else {
		fmt.Println("Client unexpectedly placed order (should have failed).")
	}
}
