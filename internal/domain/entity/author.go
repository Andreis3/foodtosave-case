package entity

import "github.com/andreis3/foodtosave-case/internal/domain/errors"

type Author struct {
	ID          string
	Name        string
	Nationality string
	errors.NotificationErrors
}

func (a *Author) Validate() *errors.NotificationErrors {
	if a.Name == "" {
		a.AddErrors(`name: is required`)
	}
	if a.Nationality == "" {
		a.AddErrors(`nationality: is required`)
	}
	return &a.NotificationErrors
}
