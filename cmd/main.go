package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/web"
	"github.com/web/api"
	"github.com/web/persistant"
	"net/http"
	"time"
)

func main() {
	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	clock := persistant.Clock{Time: time.Time{}}

	cr := persistant.NewContactRepository(db)
	msg := persistant.NewMessageRepository(db)
	cam := persistant.NewCampaignRepository(db, clock)
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

	msgChan := make(chan web.MessageRequest)
	//workers := 5
	stopChan := make(chan struct{})
	contactWorker := web.MessageRequestWorker{
		ContactRepository: cr,
		MessageRepository: msg,
		Campaigns:         ch,
		Messages:          msgChan,
		Workers:           2,
		StopChan:          stopChan,
		FromEmail:         "n.hristov@proxiad.com",
	}

	contactWorker.Start()
	senderWorker := web.SenderWorker{
		MessageRequests: msgChan,
		Workers:         2,
		StopChan:        stopChan,
		ApiKey:          "20ba63af8ed406c6f1f569dd1fb09d23",
		SecretKey:       "f484b802912a98ac4b142e28d6d05276",
		SAPIHost:        "https://api.mailjet.com/v3.1/send",
	}
	senderWorker.Start()

	authMiddleware := web.AuthMiddleware{UserRepository: usr}
	r.Use(authMiddleware.BasicAuth)

	http.ListenAndServe(":8080", r)

}
