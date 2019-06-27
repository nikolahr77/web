package api

import (
	"encoding/json"
	"github.com/web"
	"net/http"
)

func CreateUser(cr web.UserRepository) http.HandlerFunc{
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


type UserDTO struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
	Password string `json:"password"`
	Age int `json:"age"`
	Email string `json:"email"`
	CreatedOn string `json:"created_on"`
}

type RequestUserDTO struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int `json:"age"`
}

func adaptToRequestUser(u RequestUserDTO) web.RequestUser{
	return web.RequestUser{
		Name: u.Name,
		Age: u.Age,
		Email: u.Email,
		Password: u.Password,
	}
}

func userToDTO(u web.User) UserDTO{
	return UserDTO{
		GUID: u.GUID,
		Password: u.Password,
		Name: u.Name,
		Age: u.Age,
		Email: u.Email,
		CreatedOn: u.CreatedOn,
	}
}
