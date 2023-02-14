package entities

type Stack struct {
	Url          string   `json:"link"`
	Title        string   `json:"title"`
	IsAnswered   bool     `json:"is_answered"`
	AnswerCount  uint     `json:"answer_count"`
	Score        int      `json:"score"`
	CreationDate uint64   `json:"creation_date"`
	ViewCount    uint64   `json:"view_count"`
	Tags         []string `json:"tags"`
	ClosedReason string   `json:"closed_reason"`
	// manually added
	ScoreStr        string `json:"score_str"`
	CreationDateStr string `json:"creation_date_str"`
	ViewCountStr    string `json:"view_count_str"`
}
