package web

type Message struct {
	ID		int64
	Name    string
	Content   string
}

type MessageRepository interface {
	//Get(id int64) (Message, error)
	Create(m Message) (Message, error)
	//Delete(id int64) error
	//Update(id int64, m Message) (Message, error)
}
