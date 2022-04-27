package contextx

import (
	"context"
	"fmt"
	"testing"
)

func TestContextx(t *testing.T) {
	ctx := context.Background()
	SetValueToMetadata(ctx, UserID, "1151713064@qq.com")
	str, _ := GetValueFromMetadata(ctx, UserID)
	fmt.Printf("str: %v\n", str)
}
