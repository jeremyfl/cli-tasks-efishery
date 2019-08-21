package internal

// Task structure model
type Task struct {
	ID      string `json:"_id"`
	Rev     string `json:"_rev,omitempty"`
	Content string `json:"content"`
	IsDone  bool   `json:"isDone"`
	Tag     string `json:"tag"`
}
