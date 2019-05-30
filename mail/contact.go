package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Contact struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Address string `json:"address"`
}

func GetContact (w http.ResponseWriter,r *http.Request){
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Getting contacts"))
	fmt.Println(c)
}

func CreateContact(w http.ResponseWriter,r *http.Request){
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Adding contacts"))
	fmt.Println(c)
}

func DeleteContact(w http.ResponseWriter,r *http.Request){
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Deleting contacts"))
	fmt.Println(c)
}

func EditContact(w http.ResponseWriter,r *http.Request){
	var c Contact
	json.NewDecoder(r.Body).Decode(&c)
	w.Write([]byte("Editing contacts"))
	fmt.Println(c)
}

//func SaveContactToDb(db *sql.DB,c Contact) {
//
//}

func main(){
	connStr := "user=postgres dbname=email sslmode=dissable password=1234"
	db,err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	
	r := mux.NewRouter()
	r.HandleFunc("/contacts",GetContact).Methods("GET")
	r.HandleFunc("/contacts",EditContact).Methods("PUT")
	r.HandleFunc("/contacts",CreateContact).Methods("POST")
	r.HandleFunc("/contacts",DeleteContact).Methods("DELETE")
	http.ListenAndServe(":8080",r)
	}
