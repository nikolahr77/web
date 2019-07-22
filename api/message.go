package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/web"
	"github.com/web/convert"
	"net/http"
	"time"
)

//CreateMessage decodes JSON from the request and creates a new message based on the POST request
func CreateMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m RequestMessageDTO
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		adaptedReqMsg := web.RequestMessage{}
		convert.SourceToDestination(m, &adaptedReqMsg)
		message, err := msg.Create(adaptedReqMsg)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedMsg := MessageDTO{}
		convert.SourceToDestination(message, &adaptedMsg)
		json.NewEncoder(w).Encode(adaptedMsg)
	}
}

//DeleteMessage is used to delete a message with the ID from the DELETE request
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

//UpdateMessage selects a message with ID specified in the
// request and uses the JSON from the PUT request to update the message
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
		convert.SourceToDestination(m, &adaptedReqMsg)
		message, err := msg.Update(id, adaptedReqMsg)
		if err != nil {
			http.Error(w, "Internal Error", 500)
			return
		}
		adaptedMsg := MessageDTO{}
		convert.SourceToDestination(message, &adaptedMsg)
		json.NewEncoder(w).Encode(adaptedMsg)
	}
}

//GetMessage is used to get the ID from the GET request, sends a Get request
// and returns the message with the same ID
func GetMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		message, err := msg.Get(id)
		if err != nil {
			http.Error(w, "Internal Error", 500)
			return
		}
		adaptedMsg := MessageDTO{}
		convert.SourceToDestination(message, &adaptedMsg)
		json.NewEncoder(w).Encode(adaptedMsg)
	}
}

//MessageDTO is the message database object
type MessageDTO struct {
	GUID      string
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

//RequestMessageDTO is used to return info relevant to the user
type RequestMessageDTO struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
