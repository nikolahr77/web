package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/web"
	"net/http"
	"time"
)

func CreateMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m RequestMessageDTO
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		adaptedReqMsg := web.RequestMessage{}
		SourceToDestination(m, &adaptedReqMsg)
		message, err := msg.Create(adaptedReqMsg)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedMsg := MessageDTO{}
		SourceToDestination(message, &adaptedMsg)
		json.NewEncoder(w).Encode(adaptedMsg)
	}
}

func DeleteMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := msg.Delete(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
	}
}

func UpdateMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var m RequestMessageDTO
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		adaptedReqMsg := web.RequestMessage{}
		SourceToDestination(m, &adaptedReqMsg)
		message, err := msg.Update(id, adaptedReqMsg)
		if err != nil {
			http.Error(w, "Internal Error", 500)
			return
		}
		adaptedMsg := MessageDTO{}
		SourceToDestination(message, &adaptedMsg)
		json.NewEncoder(w).Encode(adaptedMsg)
	}
}

func GetMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		message, err := msg.Get(id)
		if err != nil {
			http.Error(w, "Internal Error", 500)
			return
		}
		adaptedMsg := MessageDTO{}
		SourceToDestination(message, &adaptedMsg)
		json.NewEncoder(w).Encode(adaptedMsg)
	}
}

type MessageDTO struct {
	GUID      string
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type RequestMessageDTO struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
