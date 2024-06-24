package dto

type AuthorInput struct {
	Name        string      `json:"name"`
	Nationality string      `json:"nationality"`
	Books       []BookInput `json:"books"`
}

type BookInput struct {
	Title  string `json:"title"`
	Gender string `json:"gender"`
}
