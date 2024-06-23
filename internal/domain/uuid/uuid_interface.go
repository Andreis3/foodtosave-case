package uuid

import "github.com/google/uuid"

type IUUID interface {
	Generate() string
}
type UUID struct{}

func NewUUID() *UUID {
	return &UUID{}
}
func (u *UUID) Generate() string {
	return uuid.New().String()
}
