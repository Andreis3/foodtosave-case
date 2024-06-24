package uuidmock

import "github.com/stretchr/testify/mock"

type UUIDMock struct {
	mock.Mock
}

func (r *UUIDMock) Generate() string {
	args := r.Called()
	return args.String(0)
}
