package web

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

func BasicAuth (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
			//fmt.Println(auth)


		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		//fmt.Println(pair)
		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
		})
}

func validate(username, password string) bool {
	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
		return false
	}

	searchName := `
	SELECT password FROM users WHERE name = $1`

	usr, err := db.Query(searchName, username)
	var u user

	defer usr.Close()
		for usr.Next() {
			err := usr.Scan(&u.Password)
			if err != nil {
				fmt.Println(err)
				return false
			}
		}
	fmt.Println(u.Password)
	fmt.Println(password)
	passByte := []byte(u.Password)
	requestPassByte := []byte(password)
	err = bcrypt.CompareHashAndPassword(passByte, requestPassByte)
	if err != nil {
		log.Println(err)
		return false
	}
	if username == "nikola" && password == "1234" {
		return true
	}
	return true
}

type user struct{
	Name string `db:"name"`
	Password string `db:"password"`
}
