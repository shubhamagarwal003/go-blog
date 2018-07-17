package middlewares

import (
	"context"
	"github.com/shubhamagarwal003/blog/helper"
	"net/http"
)

func UserLogged(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("SessionId")
		if err == nil {
			user := helper.Str.GetUser(cookie.Value)
			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	}
}
