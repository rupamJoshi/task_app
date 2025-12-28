package models

type Status string

const (
	TODO        Status = "TODO"
	IN_PROGRESS Status = "IN_PROGRESS"
	DONE        Status = "DONE"
)

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      Status  `json:"status"`
}
