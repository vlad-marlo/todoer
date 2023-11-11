//go:generate easyjson -all error.go
package model

type ErrorMessage struct {
	Endpoint string `json:"endpoint"`
	Code     int    `json:"code"`
	Status   string `json:"status"`
}

func (e ErrorMessage) Error() string {
	resp, err := e.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(resp)
}
