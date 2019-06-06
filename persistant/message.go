package persistant

import (
	"database/sql"
	"github.com/web"
)

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
