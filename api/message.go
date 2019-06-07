package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/web"
	"net/http"
	"strconv"
)

type MessageDTO struct {
	Name    string `json:"name"`
	Content   string `json:"content"`
}

func adaptDTOToMessage(m MessageDTO) web.Message{
	return web.Message{
		Name: m.Name,
		Content: m.Content,
	}
}

func CreateMessage(msg web.MessageRepository) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var m MessageDTO
		json.NewDecoder(r.Body).Decode(&m)
		result := adaptDTOToMessage(m)
		_,err := msg.Create(result)
		if err != nil {
			panic(err)
		}
	}
}

func DeleteMessage(msg web.MessageRepository) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		err1 := msg.Delete(idInt)
		if err1 != nil {
			panic(err1)
		}
	}
}