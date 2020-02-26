package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/Hqqm/paygo/internal/auth/interfaces"
)

type AuthMiddleware struct {
	usecases interfaces.AuthUsecases
}

func NewAuthMiddleware(usecases interfaces.AuthUsecases) *AuthMiddleware {
	return &AuthMiddleware{usecases: usecases}
}

func (am *AuthMiddleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("x-access-token")
		header = strings.TrimSpace(header)
		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := r.Context()
		account, err := am.usecases.ParseToken(ctx, header)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		}

		ctx = context.WithValue(ctx, "account", account)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
