package contextx

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc/metadata"
)

const (
	UserID = "user_id"
)

// 从metadata中获取key对应的value
func GetValue(ctx context.Context, key string) (string, error) {
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		return md.Get(key)[0], nil
	}
	return "", errors.New("fail to get userID From IncomingContext")
}

// 获取metadata中获取所有的values
func GetValues(ctx context.Context) map[string]string {
	result := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for k, v := range md {
			result[k] = v[0]
		}
	}
	return result
}

// 添加key-value到metadata中
func SetValue(ctx context.Context, key string, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}

// 批量设置metadata中的key-value
func SetValuesBatch(ctx context.Context, values map[string]string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(values))
}

// 装载Value的middleware
func LoadValues(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if uid, ok := ctx.Value(UserID).(string); ok {
			ctx = SetValuesBatch(ctx, map[string]string{
				UserID: uid,
			})
		}
		next(rw, r)
	}
}
