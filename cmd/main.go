package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/web"
	"github.com/web/api"
	"github.com/web/persistant"
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
	//s := r.Host("/users").Subrouter()
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

	ch := make(chan web.Campaign)
	r.HandleFunc("/campaign/start/{id}", api.StartCampaign(cam, ch)).Methods("POST") //post zaqvka

	contactsChan := make(chan web.SendContacts)
	//workers := 5
	stopChan := make(chan struct{})
	contactWorker := web.ContactWorker{
		ContactRepository: cr,
		Contacts:          contactsChan,
		Campaigns:         ch,
		Workers:           2,
		StopChan:          stopChan,
	}

	contactWorker.Start()
	senderWorker := web.SenderWorker{
		ContactRepository: cr,
		MessageRepository: msg,
		Contacts:          contactsChan,
		Workers:           2,
		StopChan:          stopChan,
	}
	senderWorker.Start()

	authMiddleware := web.AuthMiddleware{UserRepository: usr}
	r.Use(authMiddleware.BasicAuth)

	http.ListenAndServe(":8080", r)

}
