package errors

import "fmt"

type NotificationErrors struct {
	notification []string
}

func NewNotificationErrors() *NotificationErrors {
	return &NotificationErrors{}
}

func (n *NotificationErrors) AddErrors(notification string) {
	n.notification = append(n.notification, notification)
}

func (n *NotificationErrors) HasErrors() bool {
	return len(n.notification) > 0
}

func (n *NotificationErrors) ReturnErrors() []string {
	return n.notification
}

func (n *NotificationErrors) MergeErrors(index int, field string, childErrors *NotificationErrors) {
	for _, err := range childErrors.notification {
		n.notification = append(n.notification, fmt.Sprintf(`%s[%d].%s`, field, index, err))
	}
}
