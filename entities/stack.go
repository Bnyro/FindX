package entities

type Stack struct {
	Link       string `json:"link"`
	Title      string `json:"title"`
	IsAnswered bool   `json:"is_answered"`
}
