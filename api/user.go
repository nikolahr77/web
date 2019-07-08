package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/web"
	"github.com/web/convert"
	"net/http"
	"time"
)

func GetUser(cr web.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		user, err := cr.Get(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedUser := UserDTO{}
		convert.SourceToDestination(user, &adaptedUser)
		json.NewEncoder(w).Encode(adaptedUser)
	}
}

func UpdateUser(cr web.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var u RequestUserDTO
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		usr := web.RequestUser{}
		convert.SourceToDestination(u, &usr)
		user, err := cr.Update(id, usr)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedUser := UserDTO{}
		convert.SourceToDestination(user, &adaptedUser)
		json.NewEncoder(w).Encode(adaptedUser)
	}
}

func CreateUser(cr web.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u RequestUserDTO
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}

		usr := web.RequestUser{}
		convert.SourceToDestination(u, &usr)
		user, err := cr.Create(usr)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedUser := UserDTO{}
		convert.SourceToDestination(user, &adaptedUser)
		json.NewEncoder(w).Encode(adaptedUser)
	}
}

func DeleteUser(cr web.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := cr.Delete(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
	}
}

type UserDTO struct {
	GUID      string    `json:"guid"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type RequestUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
