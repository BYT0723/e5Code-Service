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
	return "", errors.New("fail to get userid from incomingContext")
}
