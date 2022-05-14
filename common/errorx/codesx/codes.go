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
	DockerError
	CopierError

	NotFound
	NotVerify
	AlreadyExist
	PasswordNotMatch
	TokenGenerateError
)
