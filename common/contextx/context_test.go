package contextx

import (
	"context"
	"fmt"
	"testing"
)

func TestContextx(t *testing.T) {
	ctx := context.Background()
	ctx = SetValuesBatch(ctx, map[string]string{"name": "wangtao"})
	values := GetValues(ctx)
	fmt.Printf("values: %v\n", values)
}
