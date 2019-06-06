package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/web"
	"net/http"
	"strconv"
)

type ContactDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func adaptToDTO(c web.Contact) ContactDTO {
	return ContactDTO{
		Name:    c.Name,
		Email:   c.Email,
		Age:     c.Age,
		Address: c.Address,
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

func GetContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			http.Error(w, "Bad request", 400)
		}
		c, err := cr.Get(idInt)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		fmt.Fprint(w, c)
	}

}

func DeleteContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		err1 := cr.Delete(idInt)
		if err1 != nil {
			panic(err1)
		}
	}
}

func CreateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c ContactDTO
		json.NewDecoder(r.Body).Decode(&c)
		con := adaptDTOToContact(c)
		_, err := cr.Create(con)
		if err != nil {
			panic(err)
		}
	}
}

func UpdateContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c ContactDTO
		json.NewDecoder(r.Body).Decode(&c)
		con := adaptDTOToContact(c)
		id := mux.Vars(r)["id"]
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		_, err1 := cr.Update(idInt, con)
		if err1 != nil {
			panic(err1)
		}
	}
}

//func EditContact(w http.ResponseWriter, r *http.Request) {
//	id := mux.Vars(r)["contact_id"]
//	var c ContactDTO
//	json.NewDecoder(r.Body).Decode(&c)
//	w.Write([]byte("Editing contacts"))
//	fmt.Println(c)
//
//	connStr := "user=postgres dbname=api sslmode=disable password=1234"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Connection to DB success")
//
//
//	sqlStatement := `
//	UPDATE contacts
//	SET name=$1,email=$2,age=$3,address=$4
//	WHERE id=$5;`
//
//	_, err1 := db.Exec(sqlStatement,c.Name,c.Email,c.Age,c.Address,id)
//	if err1 != nil{
//		panic(err1)
//	}
//	fmt.Println("Contact Updated")
//}

//func CreateContact(w http.ResponseWriter, r *http.Request) {
//	var c ContactDTO
//	json.NewDecoder(r.Body).Decode(&c)
//	w.Write([]byte("Adding contacts"))
//	fmt.Println(c)
//
//	connStr := "user=postgres dbname=api sslmode=disable password=1234"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Connection to DB success")
//
//	sqlStatement := `
//	INSERT INTO contacts (name,email,age,address)
//VALUES ($1, $2, $3, $4);`
//
//	_, err1 := db.Exec(sqlStatement, c.Name, c.Email, c.Age, c.Address)
//	if err1 != nil {
//		fmt.Println(err1)
//	}
//	fmt.Print("Contact inserted into DB")
//}

//func DeleteContact(w http.ResponseWriter, r *http.Request) {
//	var c Contact
//	json.NewDecoder(r.Body).Decode(&c)
//	w.Write([]byte("Deleting contacts"))
//	fmt.Println(c)
//
//	connStr := "user=postgres dbname=api sslmode=disable password=1234"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Connection to DB success")
//
//	sqlStatement := `
//	DELETE FROM contacts WHERE name=$1;`
//
//	_, err1 := db.Exec(sqlStatement,c.Name)
//	if err1 != nil{
//		panic(err1)
//	}
//	fmt.Println("Contact Deleted from DB")
//}
//