package apperror

import "encoding/json"

var (
	BadRequest = New(nil, "bad request", "", "ERROR_BAD_REQUEST")
	NotFound   = New(nil, "not found", "", "ERROR_NOT_FOUND")
)

type Error struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func New(err error, message, developerMessage, code string) *Error {
	return &Error{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func Wrap(err error) *Error {
	return New(err, "internal system error", err.Error(), "INTERNAL")
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}
