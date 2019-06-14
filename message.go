package web

import "time"

type Message struct {
	GUID    string
	Name    string
	Content string
	CreatedOn time.Time
	UpdatedOn time.Time
}

type RequestMessage struct {
	Name    string
	Content string
} 

type MessageRepository interface {
	//Get(id string) (Message, error)
	Create(m Message) (Message, error)
	Delete(id string) error
	//Update(id string, m Message) (Message, error)
}
