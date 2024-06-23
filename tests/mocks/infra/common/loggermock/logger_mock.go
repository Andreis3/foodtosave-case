package loggermock

import "github.com/stretchr/testify/mock"

type LoggerMock struct {
	mock.Mock
}

func (l *LoggerMock) DebugJson(msg string, info ...any) {
	l.Called(msg, info)

}
func (l *LoggerMock) InfoJson(msg string, info ...any) {
	l.Called(msg, info)
}
func (l *LoggerMock) WarnJson(msg string, info ...any) {
	l.Called(msg, info)
}
func (l *LoggerMock) ErrorJson(msg string, info ...any) {
	l.Called(msg, info)
}
func (l *LoggerMock) DebugText(msg string, info ...any) {
	l.Called(msg, info)
}
func (l *LoggerMock) InfoText(msg string, info ...any) {
	l.Called(msg, info)
}
func (l *LoggerMock) WarnText(msg string, info ...any) {
	l.Called(msg, info)
}
func (l *LoggerMock) ErrorText(msg string, info ...any) {
	l.Called(msg, info)
}
