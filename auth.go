package web

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	UserRepository UserRepository
}

//BasicAuth is a middleware used to verify the user from the request
func (am AuthMiddleware) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/users" {
			next.ServeHTTP(w, r)
			return
		}
		usr, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "authentication failed", http.StatusUnauthorized)
			return
		}
		user, err := am.UserRepository.GetByName(usr)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", user.GUID)
		r = r.WithContext(ctx)
		//context.Set(r, "userID", user.GUID)
		passByte := []byte(user.Password)
		requestPassByte := []byte(pass)
		err = bcrypt.CompareHashAndPassword(passByte, requestPassByte)
		if err != nil {
			log.Println(err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
