package web

import "time"

type Message struct {
	GUID      string
	Name      string
	Content   string
	CreatedOn time.Time
	UpdatedOn time.Time
	UserID    string
}

type RequestMessage struct {
	Name    string
	Content string
	UserID  string
}

type MessageRepository interface {
	Get(id string) (Message, error)
	Create(m RequestMessage, userID string) (Message, error)
	Delete(id string, userID string) error
	Update(id string, m RequestMessage, userID string) (Message, error)
}
