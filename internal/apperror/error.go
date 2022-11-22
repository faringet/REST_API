package apperror

import "encoding/json"

// создаем коды ошибок (по коду будет понимание что за ошибка, чья, какого сервиса)
var (
	ErrNotFound = NewAppError(nil, "not found", "", "US-000003")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (e *AppError) Error() string { // метод error должен соответсвовать интерфейсу error
	return e.Message
}

func (e *AppError) Unwrap() error { //соосветсвуем wrapper
	return e.Err
}

func (e *AppError) Marshal() []byte { //маршалим ошибку
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func systemError(err error) *AppError { //метод оборачаивает "не мои" ошибки в системные
	return NewAppError(err, "internal system error", err.Error(), "US-000000")

}
