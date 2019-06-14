package persistant

import (
	"database/sql"
	"github.com/web"
	"time"
)

type messageRepository struct {
	db *sql.DB
}

type messageEntity struct {
	GUID    string  `db:"guid"`
	Name    string `db:"name"`
	Content string `db:"content"`
	CreatedOn time.Time `db:"created_on"`
	UpdatedOn time.Time `db:"updated_on"`
}

func NewMessageRepository(db *sql.DB) web.MessageRepository {
	return messageRepository{db: db}
}

func adaptToMessage(m messageEntity) web.Message {
	return web.Message{
		GUID:    m.GUID,
		Name:    m.Name,
		Content: m.Content,
	}
}

func (m messageRepository) Create(msg web.Message) (web.Message, error) {
	query := `
	INSERT INTO messages (name, content)
	VALUES ($1,$2);`
	_, err := m.db.Exec(query, msg.Name, msg.Content)
	return msg, err
}

func (m messageRepository) Delete(id string) error {
	query := `
	DELETE FROM messages WHERE id=$1`
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (m messageRepository) Update(id string, msg web.Message) (web.Message, error) {
	query := `
	UPDATE messages 
	SET name=$1, content=$2
	WHERE id=$3`
	_, err := m.db.Exec(query, msg.Name, msg.Content, id)
	return msg, err
}

func (m messageRepository) Get(id string) (web.Message, error) {
	query := `
	SELECT * FROM messages WHERE id=$1`

	var e messageEntity

	rows, err := m.db.Query(query, id)
	if err != nil {
		return web.Message{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.ID, &e.Name, &e.Content)
		if err != nil {
			return web.Message{}, err
		}
	}
	result := adaptToMessage(e)
	return result, err
}
