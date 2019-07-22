package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/web"
	"github.com/web/convert"
	"net/http"
	"time"
)

//GetUser is used to get the ID from the GET request, sends a Get request
// and returns the user with the same ID
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

//UpdateUser selects a user with ID specified in the
// request and uses the JSON from the PUT request to update the user
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

//CreateUser decodes JSON from the request and creates a new user based on the POST request
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

//DeleteUser is used to delete a user with the ID from the DELETE request
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

//UserDTO is the user database object
type UserDTO struct {
	GUID      string    `json:"guid"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

//RequestUserDTO is used to return info relevant to the user
type RequestUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
