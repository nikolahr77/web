package persistant

import (
	"database/sql"
	"github.com/web"
)

type contactRepository struct {
	db *sql.DB
}

type contactEntity struct {
	ID int64 `db:"id"`
	Name string `db:"name"`
	Email   string `db:"email"`
	Age     int `db:"age"`
	Address string `db:"address"`
}

func (c contactRepository) Get(id int64) (web.Contact,error) {
	query := `SELECT name, email, age, address FROM contacts;`
	var e contactEntity
	c.db.Exec(query, e)
}

func (c contactRepository) Post(id int64) (web.Contact,error)  {
	query := `
	INSERT INTO contacts (name,email,age,address)
	VALUES ($1, $2, $3, $4);`
	var e contactEntity
	c.db.Exec(query,)
}

func (c contactRepository) Delete(id int64) (web.Contact,error)  {
	query := `
	DELETE FROM contacts WHERE name=$1;`
	var e contactEntity
	c.db.Exec(query,e)
}

func (c contactRepository)  Put{

}


func NewContactRepository(db *sql.DB) web.ContactRepository  {
	return contactRepository{db:db}
}

