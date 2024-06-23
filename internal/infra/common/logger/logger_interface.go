package logger

type ILogger interface {
	DebugJson(msg string, info ...any)
	InfoJson(msg string, info ...any)
	WarnJson(msg string, info ...any)
	ErrorJson(msg string, info ...any)
	DebugText(msg string, info ...any)
	InfoText(msg string, info ...any)
	WarnText(msg string, info ...any)
	ErrorText(msg string, info ...any)
}
