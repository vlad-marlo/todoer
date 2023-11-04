package model

type (
	// CreateTaskRequest creates task to storage
	CreateTaskRequest struct {
		Value string `json:"value"`
	}
	// GetTasksResponse return slice of tasks and some other
	// additional information.
	GetTasksResponse struct {
		Count    int       `json:"count"`
		Next     string    `json:"next"`
		Previous string    `json:"previous"`
		Tasks    []TaskDTO `json:"tasks"`
	}
	// CreateTaskResponse is returned when task was successfuly created.
	CreateTaskResponse TaskDTO
)
