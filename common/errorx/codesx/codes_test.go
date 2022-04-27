package codesx

import (
	"fmt"
	"testing"
)

func TestCodes(t *testing.T) {
	fmt.Printf("DefaultCode: %v\n", DefaultCode)
	fmt.Printf("SQLError: %v\n", SQLError)
	fmt.Printf("JSONMarshalError: %v\n", JSONMarshalError)
	fmt.Printf("ContextError: %v\n", ContextError)
	fmt.Printf("GitError: %v\n", GitError)
	fmt.Printf("DockeError: %v\n", DockerError)
	fmt.Printf("NotFound: %v\n", NotFound)
	fmt.Printf("AlreadyExist: %v\n", AlreadyExist)
	fmt.Printf("PasswordNotMatch: %v\n", PasswordNotMatch)
	fmt.Printf("TokenGenerateError: %v\n", TokenGenerateError)
}
