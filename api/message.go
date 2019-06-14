package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/web"
	"net/http"
	"time"
)

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

func CreateMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m RequestMessageDTO
		json.NewDecoder(r.Body).Decode(&m)
		result := adaptToRequestMessage(m)
		message, err := msg.Create(result)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func DeleteMessage(msg web.MessageRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err1 := msg.Delete(id)
		if err1 != nil {
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
			fmt.Println(err)
			http.Error(w, "Internal Error", 500)
		}
		json.NewEncoder(w).Encode(message)
	}
}

//func GetMessage(msg web.MessageRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := mux.Vars(r)["id"]
//		idInt, err := strconv.ParseInt(id, 10, 64)
//		if err != nil {
//			panic(err)
//		}
//		f, err1 := msg.Get(idInt)
//		if err1 != nil {
//			panic(err1)
//		}
//		fmt.Fprint(w, f)
//	}
//}
