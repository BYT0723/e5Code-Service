package codesx

const (
	RPCError = 300
	APIError = 400
)

const (
	DefaultCode = 500 + iota
	SQLError
	JSONMarshalError
	ContextError
	GitError
	DockeError

	NotFound
	AlreadyExist
	PasswordNotMatch
	TokenGenerateError
)
