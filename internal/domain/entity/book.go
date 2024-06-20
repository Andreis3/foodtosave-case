package entity

import "github.com/andreis3/foodtosave-case/internal/domain/notification"

type Book struct {
	ID     string
	Title  string
	Gender string
}

func (b *Book) Validate() *notification.Error {
	err := notification.NewError()
	if b.Title == "" {
		err.AddErrors(`title: is required`)
	}
	if b.Gender == "" {
		err.AddErrors(`gender: is required`)
	}
	return err
}
