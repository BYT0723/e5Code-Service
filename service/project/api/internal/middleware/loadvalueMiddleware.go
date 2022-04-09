package middleware

import (
	"e5Code-Service/common/contextx"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadValueMiddleware struct {
}

func NewLoadValueMiddleware() *LoadValueMiddleware {
	return &LoadValueMiddleware{}
}

func (m *LoadValueMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userid, err := contextx.GetValueFromContext(ctx, contextx.UserID)
		if err != nil {
			logx.Error("Fail to get value from context when LoadValue:", err.Error())
		}
		ctx = contextx.SetValueToMetadata(ctx, contextx.UserID, userid)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
