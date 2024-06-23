package dto

type AuthorOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Books       []struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Gender string `json:"gender"`
	} `json:"books"`
}
