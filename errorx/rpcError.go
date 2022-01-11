package errorx

import "fmt"

const (
	ExecSQLError = "ExecSQLError"

	// Service ERROR
	UserNotFound       = "UserNotFound"
	PasswordNoMatch    = "PasswordNoMatch"
	TokenGenerateError = "TokenGenerateError"
)

type RpcError struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func NewRpcError(errorType, msg string) error {
	return &RpcError{Type: errorType, Msg: msg}
}

func (e *RpcError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Msg)
}
