package common

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	id := GenUUID()
	fmt.Printf("len(id): %v\n", len(id))
}
