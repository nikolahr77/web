package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/web"
	"github.com/web/api"
	"github.com/web/persistant"
	"log"
	"net/http"
)

func main() {




	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	cr := persistant.NewContactRepository(db)
	msg := persistant.NewMessageRepository(db)
	cam := persistant.NewCampaignRepository(db)
	usr := persistant.NewUserRepository(db)
	r := mux.NewRouter()

	//finalHandler := http.HandlerFunc(api.GetContact(cr))


	r.HandleFunc("/contacts/{id}", api.GetContact(cr)).Methods("GET")
	r.HandleFunc("/contacts/{id}", api.UpdateContact(cr)).Methods("PUT")
	r.HandleFunc("/contacts", api.CreateContact(cr)).Methods("POST")
	r.HandleFunc("/contacts/{id}", api.DeleteContact(cr)).Methods("DELETE")

	r.HandleFunc("/messages/{id}", api.GetMessage(msg)).Methods("GET")
	r.HandleFunc("/messages/{id}", api.UpdateMessage(msg)).Methods("PUT")
	r.HandleFunc("/messages", api.CreateMessage(msg)).Methods("POST")
	r.HandleFunc("/messages/{id}", api.DeleteMessage(msg)).Methods("DELETE")

	r.HandleFunc("/campaign/{id}", api.GetCampaign(cam)).Methods("GET")
	r.HandleFunc("/campaign/{id}", api.UpdateCampaign(cam)).Methods("PUT")
	r.HandleFunc("/campaign", api.CreateCampaign(cam)).Methods("POST")
	r.HandleFunc("/campaign/{id}", api.DeleteCampaign(cam)).Methods("DELETE")

	r.HandleFunc("/users/{id}", api.GetUser(usr)).Methods("GET")
	r.HandleFunc("/users/{id}", api.UpdateUser(usr)).Methods("PUT")
	r.HandleFunc("/users", api.CreateUser(usr)).Methods("POST") //bez middleware
	r.HandleFunc("/users/{id}", api.DeleteUser(usr)).Methods("DELETE")
	//http.Handle("/campaign/{id}", middlewareOne(r))
	r.Use(web.BasicAuth)


	http.ListenAndServe(":8080", r)



}


func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		w.Write([]byte("OK"))
		next.ServeHTTP(w, r)
		fmt.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		//if r.URL.Path != "/contacts/{id}" {
		//	return
		//}
		next.ServeHTTP(w, r)
		fmt.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}
