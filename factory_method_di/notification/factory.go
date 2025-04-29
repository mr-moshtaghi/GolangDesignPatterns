package notification

import "fmt"

type NotificationCreator interface {
	Create(notificationType string) (Notification, error)
}

type NotificationFactory struct {
	creators map[string]func() Notification
}

func NewNotificationFactory() *NotificationFactory {
	return &NotificationFactory{
		creators: make(map[string]func() Notification),
	}
}

func (nf *NotificationFactory) Register(notificationType string, creator func() Notification) {
	nf.creators[notificationType] = creator
}

func (nf *NotificationFactory) Create(notificationType string) (Notification, error) {
	creator, exists := nf.creators[notificationType]
	if !exists {
		return nil, fmt.Errorf("notificatio %s don't support", notificationType)
	}
	return creator(), nil
}
