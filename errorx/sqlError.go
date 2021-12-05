package errorx

const (
	// code error
	DefaultSQLError = 2000
	UserNotFound    = 2001
	PasswordNoMatch = 2002
)

type SQLError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SQLErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewSQLError(code int, msg string) error {
	return &SQLError{Code: code, Msg: msg}
}

func NewDefaultSQLError() error {
	return &SQLError{
		Code: DefaultSQLError,
		Msg:  "default sql error",
	}
}

func (e *SQLError) Error() string {
	return e.Msg
}

func (e *SQLError) Data() *SQLErrorResponse {
	return &SQLErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
