package common

type TaskStatus int

const (
	TaskStatusOpen       TaskStatus = 0
	TaskStatusInProgress TaskStatus = 1
	TaskStatusCompleted  TaskStatus = 2
)

type Task struct {
	ID      int        `json:"id"`
	Status  TaskStatus `json:"status"`
	Message string     `json:"message"`
}

const UnsetTaskID int = -1

type Tasks struct {
	Items []Task `json:"items"`
}
