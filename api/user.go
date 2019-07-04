package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/web"
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
		json.NewEncoder(w).Encode(userToDTO(user))
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
		usr := adaptToRequestUser(u)
		user, err := cr.Update(id, usr)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		json.NewEncoder(w).Encode(userToDTO(user))
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
		usr := adaptToRequestUser(u)
		user, err := cr.Create(usr)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		json.NewEncoder(w).Encode(userToDTO(user))
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

func adaptToRequestUser(u RequestUserDTO) web.RequestUser {
	return web.RequestUser{
		Name:     u.Name,
		Age:      u.Age,
		Email:    u.Email,
		Password: u.Password,
	}
}

func userToDTO(u web.User) UserDTO {
	return UserDTO{
		GUID:      u.GUID,
		Password:  u.Password,
		Name:      u.Name,
		Age:       u.Age,
		Email:     u.Email,
		CreatedOn: u.CreatedOn,
		UpdatedOn: u.UpdatedOn,
	}
}
