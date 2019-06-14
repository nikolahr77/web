package persistant

import (
	"database/sql"
	"github.com/google/uuid"
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
	INSERT INTO messages (guid, name, content, created_on, updated_on)
	VALUES ($1,$2,$3,$4,$5);`
	uuid := uuid.New()
	createdOn := time.Now().UTC()
	_, err := m.db.Exec(query, uuid, msg.Name, msg.Content,createdOn,createdOn)
	return web.Message{
			GUID: uuid.String(),
			Name: msg.Name,
			Content: msg.Content,
			CreatedOn: createdOn,
			UpdatedOn: createdOn,
	}, err
}

func (m messageRepository) Delete(id string) error {
	query := `
	DELETE FROM messages WHERE guid=$1`
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

//func (m messageRepository) Update(id string, msg web.Message) (web.Message, error) {
//	query := `
//	UPDATE messages
//	SET name=$1, content=$2
//	WHERE id=$3`
//	_, err := m.db.Exec(query, msg.Name, msg.Content, id)
//	return msg, err
//}
//
//func (m messageRepository) Get(id string) (web.Message, error) {
//	query := `
//	SELECT * FROM messages WHERE id=$1`
//
//	var e messageEntity
//
//	rows, err := m.db.Query(query, id)
//	if err != nil {
//		return web.Message{}, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&e.ID, &e.Name, &e.Content)
//		if err != nil {
//			return web.Message{}, err
//		}
//	}
//	result := adaptToMessage(e)
//	return result, err
//}
