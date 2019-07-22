package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"github.com/web/convert"
	"time"
)

//Create adds a new message to the DB
func (m messageRepository) Create(msg web.RequestMessage) (web.Message, error) {
	query := `
	INSERT INTO messages (guid, name, content, created_on, updated_on)
	VALUES ($1,$2,$3,$4,$5);`
	uuid := uuid.New()
	createdOn := time.Now().UTC()
	_, err := m.db.Exec(query, uuid, msg.Name, msg.Content, createdOn, createdOn)
	return web.Message{
		GUID:      uuid.String(),
		Name:      msg.Name,
		Content:   msg.Content,
		CreatedOn: createdOn,
		UpdatedOn: createdOn,
	}, err
}

//Delete is used to remove a message from the DB by a given ID.
func (m messageRepository) Delete(id string) error {
	query := `
	DELETE FROM messages WHERE guid=$1`
	_, err := m.db.Exec(query, id)
	return err
}

//Update searches the DB for a message by a given
// ID and updates the message with the given RequestMessage
func (m messageRepository) Update(id string, msg web.RequestMessage) (web.Message, error) {
	query := `
	UPDATE messages
	SET name=$1, content=$2, updated_on=$3
	WHERE guid=$4`
	updatedOn := time.Now().UTC()
	_, err := m.db.Exec(query, msg.Name, msg.Content, updatedOn, id)
	return web.Message{
		Name:      msg.Name,
		Content:   msg.Content,
		UpdatedOn: updatedOn,
	}, err
}

//Get is used to return a message from the DB by a given ID.
func (m messageRepository) Get(id string) (web.Message, error) {
	query := `
	SELECT * FROM messages WHERE guid=$1`

	var e messageEntity
	rows, err := m.db.Query(query, id)
	if err != nil {
		return web.Message{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.GUID, &e.Name, &e.Content, &e.CreatedOn, &e.UpdatedOn)
		if err != nil {
			return web.Message{}, err
		}
	}
	result := web.Message{}
	convert.SourceToDestination(e, &result)
	return result, err
}

type messageRepository struct {
	db *sql.DB
}

type messageEntity struct {
	GUID      string    `db:"guid"`
	Name      string    `db:"name"`
	Content   string    `db:"content"`
	CreatedOn time.Time `db:"created_on"`
	UpdatedOn time.Time `db:"updated_on"`
}

func NewMessageRepository(db *sql.DB) web.MessageRepository {
	return messageRepository{db: db}
}
