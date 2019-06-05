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
		Name: c.Name,
		Email: c.Email,
		Age: c.Age,
		Address: c.Address,
	}
}


func GetContact(cr web.ContactRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		idInt, err := strconv.ParseInt(id,10,64)
		if err != nil {
			panic(err)
		}
		c, err := cr.Get(idInt)
		dto := adaptToDTO(c)
		json.NewEncoder(w).Encode(dto)

		for c.Next() {
			err := rows.Scan(&e.Name, &e.Email, &e.Age, &e.Address)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, e.Name, e.Email, e.Age, e.Address)
		}
	}
	}

}

//func DeleteContact (cr web.ContactRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := mux.Vars(r)["contact_id"]
//		web.ContactRepository(DeleteContact(id))
//	}
//}
//

//func DTOToAdapt(c web.Contact) ContactDTO {
//	return ContactDTO{
//		c.Name: Name,
//		Email: c.Email,
//		Age: c.Age,
//		Address: c.Address,
//	}
//}
//
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
