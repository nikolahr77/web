package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"github.com/web/convert"
	"time"
)

//Create adds a new message to the DB
func (m messageRepository) Create(msg web.RequestMessage, userID string) (web.Message, error) {
	query := `
	INSERT INTO messages (guid, name, content, created_on, updated_on, userID)
	VALUES ($1,$2,$3,$4,$5,$6);`
	uuid := uuid.New()
	createdOn := time.Now().UTC()
	_, err := m.db.Exec(query, uuid, msg.Name, msg.Content, createdOn, createdOn, userID)
	return web.Message{
		GUID:      uuid.String(),
		Name:      msg.Name,
		Content:   msg.Content,
		CreatedOn: createdOn,
		UpdatedOn: createdOn,
		UserID:    userID,
	}, err
}

//Delete is used to remove a message from the DB by a given ID.
func (m messageRepository) Delete(id string, userID string) error {
	query := `
	DELETE FROM messages WHERE guid=$1 AND userID = $2`
	_, err := m.db.Exec(query, id, userID)
	return err
}

//Update searches the DB for a message by a given
// ID and updates the message with the given RequestMessage
func (m messageRepository) Update(id string, msg web.RequestMessage, userID string) (web.Message, error) {
	query := `
	UPDATE messages
	SET name=$1, content=$2, updated_on=$3
	WHERE guid=$4 AND userID = $5`
	updatedOn := time.Now().UTC()
	_, err := m.db.Exec(query, msg.Name, msg.Content, updatedOn, id, userID)
	return web.Message{
		Name:      msg.Name,
		Content:   msg.Content,
		UpdatedOn: updatedOn,
	}, err
}

//Get is used to return a message from the DB by a given ID.
func (m messageRepository) Get(id string, userID string) (web.Message, error) {
	query := `
	SELECT * FROM messages WHERE guid=$1 AND userid=$2`

	var e messageEntity
	rows, err := m.db.Query(query, id, userID)
	if err != nil {
		return web.Message{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.GUID, &e.Name, &e.Content, &e.CreatedOn, &e.UpdatedOn, &e.UserID)
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
	UserID    string    `db: "userID"`
}

func NewMessageRepository(db *sql.DB) web.MessageRepository {
	return messageRepository{db: db}
}
