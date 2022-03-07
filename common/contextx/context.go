package contextx

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc/metadata"
)

const (
	UserID = "userID"
)

func GetValue(ctx context.Context, key string) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md.Get(key)[0], nil
	}
	return "", errors.New("fail to get userid from incomingContext")
}

func GetValues(ctx context.Context) map[string]string {
	result := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {

		for k, v := range md {
			result[k] = v[0]
		}
	} else {
		return nil
	}
	return result
}

func SetValue(ctx context.Context, key string, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}

func SetValuesBatch(ctx context.Context, values map[string]string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(values))
}

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
