package persistant

import (
	"database/sql"
	"github.com/web"
)

type messageRepository struct {
	db *sql.DB
}

type messageEntity struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Content   string `db:"content"`
}

func NewMessageRepository(db *sql.DB) web.MessageRepository {
	return messageRepository{db: db}
}

func (m messageRepository) Create(msg web.Message) (web.Message, error){
	query := `
	INSERT INTO messages (name, content)
	VALUES ($1,$2);`
	_, err := m.db.Exec(query,msg.Name,msg.Content)
	return msg,err
}