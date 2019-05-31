package main

import "C"
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
	w.Write([]byte("Getting contacts"))

	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)

	rows, err := db.Query(`SELECT name, email, age, address FROM contacts;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.Name, &c.Email, &c.Age, &c.Address)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, c.Name, c.Email, c.Age, c.Address)
	}
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
	fmt.Println("Connection to DB success")

	sqlStatement := `
	INSERT INTO contacts (name,email,age,address)
VALUES ($1, $2, $3, $4);`

	_, err1 := db.Exec(sqlStatement, c.Name, c.Email, c.Age, c.Address)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Print("Contact inserted into DB")
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Deleting contacts"))
	fmt.Println(c)

	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to DB success")

	sqlStatement := `
	DELETE FROM contacts WHERE name=$1;`

	_, err1 := db.Exec(sqlStatement,c.Name)
	if err1 != nil{
		panic(err1)
	}
	fmt.Println("Contact Deleted from DB")
}

func EditContact(w http.ResponseWriter, r *http.Request) {
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Editing contacts"))
	fmt.Println(c)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/contacts", GetContact).Methods("GET")
	r.HandleFunc("/contacts", EditContact).Methods("PUT")
	r.HandleFunc("/contacts", CreateContact).Methods("POST")
	r.HandleFunc("/contacts", DeleteContact).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
