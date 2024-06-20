package util

type ValidationError struct {
	Code        string
	Status      int
	Origin      string
	ClientError []string
	LogError    []string
}
