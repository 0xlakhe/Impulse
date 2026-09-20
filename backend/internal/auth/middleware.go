package auth

import (
	"net/http"
	"strings"

	"github.com/0xlakhe/Impluse/internal/httpx"
)

type Middleware struct {
	jwt *JWTManager
}

func NewMiddleware(jwt *JWTManager) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				httpx.Error(w, http.StatusUnauthorized, "missing authorization header")
				return
			}
			const prefix = "Bearer "
			if !strings.HasPrefix(authHeader, prefix) {
				httpx.Error(w, http.StatusUnauthorized, "invalid authorization header")
				return
			}
			tokenString := strings.TrimPrefix(authHeader, prefix)
			claims, err := m.jwt.Parse(tokenString)
			if err != nil {
				httpx.Error(w, http.StatusUnauthorized, err.Error())
				return
			}
			ctx := WithUserID(r.Context(), claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
