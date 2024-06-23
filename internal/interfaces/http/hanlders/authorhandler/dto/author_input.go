package dto

type AuthorInput struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Books       []struct {
		Title  string `json:"title"`
		Gender string `json:"gender"`
	} `json:"books"`
}
