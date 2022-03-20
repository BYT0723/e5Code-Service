package contextx

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

const (
	UserID = "user_id"
)

// 从metadata中获取key对应的value
func GetValueFromMetadata(ctx context.Context, key string) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md.Get(key)[0], nil
	}
	return "", errors.New("fail to get value From IncomingContext")
}

// 从Context中获取key对应的value
func GetValueFromContext(ctx context.Context, key string) (string, error) {
	if v, ok := ctx.Value(key).(string); ok {
		return v, nil
	}
	return "", errors.New("fail to get value From Context")
}

// 添加key-value到metadata中
func SetValueToMetadata(ctx context.Context, key string, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}
