package persistant

import (
	"database/sql"
	"github.com/web"
)

type contactRepository struct {
	db *sql.DB
}

type contactEntity struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Email   string `db:"email"`
	Age     int    `db:"age"`
	Address string `db:"address"`
}

func NewContactRepository(db *sql.DB) web.ContactRepository {
	return contactRepository{db: db}
}

func (c contactRepository) Get(id int64) (web.Contact, error) {
	query := `SELECT * FROM contacts WHERE id=$1;`
	var e contactEntity
	rows, err := c.db.Query(query, id)
	if err != nil {
		return web.Contact{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Address)
		if err != nil {
			return web.Contact{}, err
		}
	}
	result := adaptToContact(e)
	return result, err
}

func adaptToContact(entity contactEntity) web.Contact {
	return web.Contact{
		ID: entity.ID,
		Name:    entity.Name,
		Email:   entity.Email,
		Age:     entity.Age,
		Address: entity.Address,
	}
}

func (c contactRepository) Create(con web.Contact) (web.Contact, error) {
	query := `
	INSERT INTO contacts (name,email,age,address)
	VALUES ($1, $2, $3, $4);`
	_, err := c.db.Exec(query, con.Name, con.Email, con.Age, con.Address)
	return con, err
}

func (c contactRepository) Delete(id int64) error {
	query := `
	DELETE FROM contacts WHERE id=$1;`
	//var e contactEntity
	_, err := c.db.Exec(query, id)
	return err
}

func (c contactRepository) Update(id int64, con web.Contact) (web.Contact, error) {
	query := `
	UPDATE contacts
	SET name=$1,email=$2,age=$3,address=$4
	WHERE id=$5;`
	_, err := c.db.Exec(query, con.Name, con.Email, con.Age, con.Address, id)
	return con, err
}
