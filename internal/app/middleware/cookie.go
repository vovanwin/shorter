package middleware

import (
	"context"
	"github.com/vovanwin/shorter/internal/app/helper"
	"net/http"
)

const key = "user"

func UserCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := r.Cookie("token")
		token := new(helper.Token)
		var userUUID []byte
		if err != nil {

			userUUID, _ = token.CreateUserID()

			userCookie, err := token.Encode(userUUID)
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
			userUUID, err = token.Decode(user.Value)
			if err != nil {
				userUUID, _ = token.CreateUserID()

				userCookie, err := token.Encode(userUUID)
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

		ctx := context.WithValue(r.Context(), key, userUUID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
