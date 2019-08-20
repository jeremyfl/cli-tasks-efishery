package internal

// Task structure model
type Task struct {
	Content string `json:"content"`
	IsDone  bool   `json:"isDone"`
	Tag     string `json:"tag"`
}
