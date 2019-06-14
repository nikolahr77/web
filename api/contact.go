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

type ContactDTO struct {
	GUID      string
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

// RequestContactDTO, ContactDTO, |RequestContact, Contact,| ContactEntity (messages and campaign the same)
//         API/http               |  domain                |  DB, persistent

type RequestContactDTO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

func adaptToDTO(c web.Contact) ContactDTO {
	return ContactDTO{
		GUID: 	   c.GUID,
		Name:      c.Name,
		Email:     c.Email,
		Age:       c.Age,
		Address:   c.Address,
		CreatedOn: c.CreatedOn,
		UpdatedOn: c.UpdatedOn,
	}
}

func adaptDTOToContact(c ContactDTO) web.Contact {
	return web.Contact{
		Name:    c.Name,
		Email:   c.Email,
		Age:     c.Age,
		Address: c.Address,
	}
}

func adaptToRequestContact(c RequestContactDTO) web.RequestContact{
     return web.RequestContact{
     	Name: c.Name,
     	Age: c.Age,
     	Address: c.Address,
     	Email: c.Email,
	 }
}

func GetContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		guid := mux.Vars(r)["id"]
		c, err := cr.Get(guid)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal error", 500)
			return
		}
		fmt.Fprint(w, c)
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
		json.NewDecoder(r.Body).Decode(&c)
		con := adaptToRequestContact(c)
		contact, err := cr.Create(con)
		if err != nil {
			http.Error(w, "Internal error", 500)
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
		_, err1 := cr.Update(id, con)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
		}
	}
}
