package contextx

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

const (
	UserID = "userID"
)

func GetUserID(ctx context.Context) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md.Get(UserID)[0], nil
	}

	if id, ok := ctx.Value(UserID).(string); ok {
		return id, nil
	}
	return "", errors.New("fail to get userid from incomingContext")
}

func SetUserID(ctx context.Context, userID string) context.Context {
	md := metadata.Pairs(UserID, userID)
	return metadata.NewOutgoingContext(ctx, md)
}
