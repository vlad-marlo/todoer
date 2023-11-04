package model

type (
	CreateTaskRequest struct {
		Value string `json:"value"`
	}
	GetTasksResponse struct {
		Count    int       `json:"count"`
		Next     string    `json:"next"`
		Previous string    `json:"previous"`
		Tasks    []TaskDTO `json:"tasks"`
	}
	CreateTaskResponse TaskDTO
)
