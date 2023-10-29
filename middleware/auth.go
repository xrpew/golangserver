package middleware

import (
	"main/models"
	"main/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	NO_AUTH_NEEDED = []string{
		"/login",
		"/signup",
		"/",
	}
)

func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := r.Header.Get("Authorization")
			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTsecret), nil
			})

			if err != nil {
				http.Error(w, "Token invalido", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})

	}
}
