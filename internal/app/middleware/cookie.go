package middleware

import (
	"context"
	"github.com/vovanwin/shorter/internal/app/helper"
	"net/http"
)

func UserCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := r.Cookie("token")
		token := new(helper.Token)
		var userUuid []byte
		if err != nil {

			userUuid, _ = token.CreateUserId()

			userCookie, err := token.Encode(userUuid)
			if err != nil {
				return
			}
			cookie := &http.Cookie{
				Name:     "token",
				Value:    userCookie,
				MaxAge:   0,
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		} else {
			userUuid, err = token.Decode(user.Value)
			if err != nil {
				userUuid, _ = token.CreateUserId()

				userCookie, err := token.Encode(userUuid)
				if err != nil {
					return
				}
				cookie := &http.Cookie{
					Name:     "token",
					Value:    userCookie,
					MaxAge:   0,
					HttpOnly: true,
				}
				http.SetCookie(w, cookie)
			}
		}

		ctx := context.WithValue(r.Context(), "user", userUuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
