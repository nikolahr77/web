package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Getting contacts"))
	fmt.Println(c)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Adding contacts"))
	fmt.Println(c)


	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection success")

	sqlStatement := `
	INSERT INTO contacts (name,email,age,address)
VALUES ($1, $2, $3, $4);`

	_, err1 := db.Exec(sqlStatement,c.Name,c.Email,c.Age,c.Address)
	if err1 != nil{
		fmt.Println(err1)
	}
	fmt.Print("Contact inserted into DB")
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Deleting contacts"))
	fmt.Println(c)
}

func EditContact(w http.ResponseWriter, r *http.Request) {
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Editing contacts"))
	fmt.Println(c)
}

func SaveContactToDb(db *sql.DB,c Contact) {
	err := db.QueryRow(`INSERT INTO contacts(name, email, age, address)
	VALUES(c.Name, c.Email, c.Age, 'c.Address)`)
	if err!=nil{
		panic(err)
	}
	fmt.Print("Contact inserted into DB")
}

func main() {


	r := mux.NewRouter()
	r.HandleFunc("/contacts", GetContact).Methods("GET")
	r.HandleFunc("/contacts", EditContact).Methods("PUT")
	r.HandleFunc("/contacts", CreateContact).Methods("POST")
	r.HandleFunc("/contacts", DeleteContact).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}


