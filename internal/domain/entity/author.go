package entity

import "github.com/andreis3/foodtosave-case/internal/domain/notification"

type Author struct {
	ID          string
	Name        string
	Nationality string
	notification.Error
}

func (a *Author) Validate() *notification.Error {
	if a.Name == "" {
		a.AddErrors(`name: is required`)
	}
	if a.Nationality == "" {
		a.AddErrors(`nationality: is required`)
	}
	return &a.Error
}
