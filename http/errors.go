package http

import (
	"encoding/json"
	"time"
)

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(ErrorMessage error) *ErrorDTO {
	return &ErrorDTO{
		Message: ErrorMessage.Error(),
		Time:    time.Now(),
	}
}
func (e ErrorDTO) ErrorReadable() string {
	b, _ := json.Marshal(e.Message)
	return string(b)
}
