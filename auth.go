package web

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	UserRepository UserRepository
}

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
