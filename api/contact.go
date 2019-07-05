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
		contact, err := cr.Get(guid)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal error", 500)
			return
		}
		contactDTO := ContactDTO{}
		SourceToDestination(contact, &contactDTO)
		json.NewEncoder(w).Encode(contactDTO)
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
		con := web.RequestContact{}
		SourceToDestination(c, &con)
		contact, err1 := cr.Create(con)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		contactDTO := ContactDTO{}
		SourceToDestination(contact, &contactDTO)
		json.NewEncoder(w).Encode(contactDTO)
	}
}

func UpdateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestContactDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		con := web.RequestContact{}
		SourceToDestination(c, &con)
		id := mux.Vars(r)["id"]
		contact, err := cr.Update(id, con)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		contactDTO := ContactDTO{}
		SourceToDestination(contact, &contactDTO)
		json.NewEncoder(w).Encode(contactDTO)
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
