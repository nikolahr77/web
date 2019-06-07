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
	Content string `db:"content"`
}

func NewMessageRepository(db *sql.DB) web.MessageRepository {
	return messageRepository{db: db}
}

func (m messageRepository) Create(msg web.Message) (web.Message, error) {
	query := `
	INSERT INTO messages (name, content)
	VALUES ($1,$2);`
	_, err := m.db.Exec(query, msg.Name, msg.Content)
	return msg, err
}

func (m messageRepository) Delete(id int64) error {
	query := `
	DELETE FROM messages WHERE id=$1`
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (m messageRepository) Update(id int64, msg web.Message) (web.Message, error) {
	query := `
	UPDATE messages 
	SET name=$1, content=$2
	WHERE id=$3`
	_, err := m.db.Exec(query, msg.Name, msg.Content, id)
	return msg, err
}
