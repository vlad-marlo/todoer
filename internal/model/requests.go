//go:generate easyjson -all requests.go

package model

type (
	// CreateTaskRequest creates task to storage
	CreateTaskRequest struct {
		Value  string `json:"task"`
		Status Status `json:"status"`
	}
	// UpdateTaskRequest is request which makes available to update task.
	UpdateTaskRequest struct {
		Task   string `json:"task"`
		Status string `json:"status"`
	}
	GetManyTasksRequest struct {
		Offset int    `query:"offset" json:"-"`
		Limit  int    `query:"limit" json:"-"`
		Task   string `query:"task" json:"-"`
		Status string `query:"status" json:"-"`
	}
)
