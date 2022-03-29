package types

import "fmt"

type InternalError struct {
	ErrorCode    int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
	ErrorDetails error  `json:"errorDetails"`
}

func (e *InternalError) Error() string {
	return e.ErrorMessage
}

func UnmarshallError(err error) error {
	return &InternalError{
		ErrorCode:    1000,
		ErrorMessage: "Unable to parse object into list",
		ErrorDetails: fmt.Errorf("unmarshall error: %w", err),
	}
}

func IndexOutOfBoundError() error {
	return &InternalError{
		ErrorCode:    1001,
		ErrorMessage: "Index referred exceeds the max bounds for array",
	}
}
