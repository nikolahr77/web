package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/web"
	"net/http"
	"time"
)

func GetContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		guid := mux.Vars(r)["id"]
		c, err := cr.Get(guid)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal error", 500)
			return
		}
		json.NewEncoder(w).Encode(adaptToDTO(c))
	}

}

func DeleteContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err1 := cr.Delete(id)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
		}
	}
}

func CreateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestContactDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		con := adaptToRequestContact(c)
		contact, err1 := cr.Create(con)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		json.NewEncoder(w).Encode(adaptToDTO(contact))
	}
}

func UpdateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestContactDTO
		json.NewDecoder(r.Body).Decode(&c)
		con := adaptToRequestContact(c)
		id := mux.Vars(r)["id"]
		contact, err1 := cr.Update(id, con)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		json.NewEncoder(w).Encode(adaptToDTO(contact))
	}
}

type ContactDTO struct {
	GUID      string
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

// RequestContactDTO, ContactDTO, |RequestContact, Contact,| ContactEntity (messages and campaign the same)
//         API/http               |  domain                |  DB, persistent

type RequestContactDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func adaptToDTO(c web.Contact) ContactDTO {
	return ContactDTO{
		GUID:      c.GUID,
		Name:      c.Name,
		Email:     c.Email,
		Age:       c.Age,
		Address:   c.Address,
		CreatedOn: c.CreatedOn,
		UpdatedOn: c.UpdatedOn,
	}
}

func adaptToRequestContact(c RequestContactDTO) web.RequestContact {
	return web.RequestContact{
		Name:    c.Name,
		Age:     c.Age,
		Address: c.Address,
		Email:   c.Email,
	}
}
