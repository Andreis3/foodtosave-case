package entity

import "github.com/andreis3/foodtosave-case/internal/domain/errors"

type Book struct {
	ID     string
	Title  string
	Gender string
}

func (b *Book) Validate() *errors.NotificationErrors {
	err := errors.NewNotificationErrors()
	if b.Title == "" {
		err.AddErrors(`title: is required`)
	}
	if b.Gender == "" {
		err.AddErrors(`gender: is required`)
	}
	return err
}
