package entities

type News struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Thumbnail   string `json:"thumbnail"`
	UploadDate  string `json:"uploadDate"`
}
