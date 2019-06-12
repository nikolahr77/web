package api

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/web"
	"net/http"
	"time"
)

type ContactDTO struct {
	GUID    string
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	CreatedOn time.Time
	UpdatedOn time.Time
}

// RequestContactDTO, ContactDTO, |RequestContact, Contact,| ContactEntity (messages and campaign the same)
//         API/http               |  domain                |  DB, persistent


type RequestContactDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	CreatedOn time.Time
	UpdatedOn time.Time
}

func adaptToDTO(c web.Contact) ContactDTO {
	return ContactDTO{
		Name:    c.Name,
		Email:   c.Email,
		Age:     c.Age,
		Address: c.Address,
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

//func GetContact(cr web.ContactRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		guid := mux.Vars(r)["id"]
//		c, err := cr.Get(guid)
//		if err != nil {
//			http.Error(w, "Internal error", 500)
//		}
//		fmt.Fprint(w, c)
//	}
//
//}
//
//func DeleteContact(cr web.ContactRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := mux.Vars(r)["id"]
//		idInt, err := strconv.ParseInt(id, 10, 64)
//		if err != nil {
//			panic(err)
//		}
//		err1 := cr.Delete(idInt)
//		if err1 != nil {
//			panic(err1)
//		}
//	}
//}

func CreateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c ContactDTO
		fmt.Println("HELLO")
		json.NewDecoder(r.Body).Decode(&c)
		con := adaptDTOToContact(c)
		contact , err := cr.Create(con)
		if err != nil {
			http.Error(w,"Internal error",500)
		}
		json.NewEncoder(w).Encode(contact)
	}
}
//
//func UpdateContact(cr web.ContactRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var c ContactDTO
//		json.NewDecoder(r.Body).Decode(&c)
//		con := adaptDTOToContact(c)
//		id := mux.Vars(r)["id"]
//		//idInt, err := strconv.ParseInt(id, 10, 64)
//		//if err != nil {
//		//	panic(err)
//		//}
//		_, err1 := cr.Update(idInt, con)
//		if err1 != nil {
//			panic(err1)
//		}
//	}
//}