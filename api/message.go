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
		json.NewDecoder(r.Body).Decode(&m)
		result := adaptToRequestMessage(m)
		message, err := msg.Create(result)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		json.NewEncoder(w).Encode(adaptMessageToDTO(message))
	}
}

func DeleteMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := msg.Delete(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
	}
}

func UpdateMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var m RequestMessageDTO
		json.NewDecoder(r.Body).Decode(&m)
		result := adaptToRequestMessage(m)
		message, err := msg.Update(id, result)
		if err != nil {
			http.Error(w, "Internal Error", 500)
		}
		json.NewEncoder(w).Encode(adaptMessageToDTO(message))
	}
}

func GetMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		f, err := msg.Get(id)
		if err != nil {
			http.Error(w, "Internal Error", 500)
		}
		json.NewEncoder(w).Encode(adaptMessageToDTO(f))
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

func adaptToRequestMessage(m RequestMessageDTO) web.RequestMessage {
	return web.RequestMessage{
		Name:    m.Name,
		Content: m.Content,
	}
}

func adaptMessageToDTO(c web.Message) MessageDTO {
	return MessageDTO{
		GUID:      c.GUID,
		Name:      c.Name,
		Content:   c.Content,
		CreatedOn: c.CreatedOn,
		UpdatedOn: c.UpdatedOn,
	}
}
