package notification

import "fmt"

type Error struct {
	notification []string
}

func NewError() *Error {
	return &Error{}
}

func (n *Error) AddErrors(notification string) {
	n.notification = append(n.notification, notification)
}

func (n *Error) HasErrors() bool {
	return len(n.notification) > 0
}

func (n *Error) ReturnErrors() []string {
	return n.notification
}

func (n *Error) MergeErrors(index int, field string, childErrors *Error) {
	for _, err := range childErrors.notification {
		n.notification = append(n.notification, fmt.Sprintf(`%s[%d].%s`, field, index, err))
	}
}
