package internal

type Task struct {
	Content string `json:"content"`
	IsDone  bool   `json:"isDone"`
	Tag     string `json:"tag"`
}
