package api

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/web"
	"net/http"
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