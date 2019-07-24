package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/web"
	"github.com/web/convert"
	"net/http"
	"time"
)

//GetContact is used to get the ID from the GET request, sends a Get request
// and returns the Contact with the same ID
func GetContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		guid := mux.Vars(r)["id"]
		userID := context.Get(r, "userID").(string)
		contact, err := cr.Get(guid, userID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal error", 500)
			return
		}
		contactDTO := ContactDTO{}
		convert.SourceToDestination(contact, &contactDTO)
		json.NewEncoder(w).Encode(contactDTO)
	}
}

//DeleteContact is used to delete a contact with the ID from the DELETE request
func DeleteContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		userID := context.Get(r, "userID").(string)
		err1 := cr.Delete(id, userID)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
		}
	}
}

//CreateCampaign decodes JSON from the request and creates a new contact based on the POST request
func CreateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestContactDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		con := web.RequestContact{}
		userID := context.Get(r, "userID").(string)
		convert.SourceToDestination(c, &con)
		contact, err1 := cr.Create(con, userID)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		contactDTO := ContactDTO{}
		convert.SourceToDestination(contact, &contactDTO)
		json.NewEncoder(w).Encode(contactDTO)
	}
}

//UpdateContact selects a Contact with ID specified in the
// request and uses the JSON from the PUT request to update the contact
func UpdateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestContactDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		con := web.RequestContact{}
		convert.SourceToDestination(c, &con)
		id := mux.Vars(r)["id"]
		userID := context.Get(r, "userID").(string)
		contact, err := cr.Update(id, con, userID)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		contactDTO := ContactDTO{}
		convert.SourceToDestination(contact, &contactDTO)
		json.NewEncoder(w).Encode(contactDTO)
	}
}

//ContactDTO is the contact database object
type ContactDTO struct {
	GUID      string
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	UserID    string    `json:"user_id"`
}

// RequestContactDTO, ContactDTO, |RequestContact, Contact,| ContactEntity (messages and campaign the same)
//         API/http               |  domain                |  DB, persistent

//RequestContactDTO is used to return info relevant to the user
type RequestContactDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	UserID  string `json:"user_id"`
}
