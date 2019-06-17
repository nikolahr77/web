package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"time"
)

func (c contactRepository) Get(id string) (web.Contact, error) {
	query := `SELECT * FROM contacts WHERE id=$1;`
	var e contactEntity
	rows, err := c.db.Query(query, id)
	if err != nil {
		return web.Contact{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Address, &e.CreatedOn, &e.UpdatedOn)
		if err != nil {
			return web.Contact{}, err
		}
	}
	result := adaptToContact(e)
	return result, err
}

func (c contactRepository) Create(con web.RequestContact) (web.Contact, error) {
	uuid := uuid.New()
	query := `
	INSERT INTO contacts (id, name,email,age,address,created_on,updated_on)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`
	createdOn := time.Now().UTC()
	_, err := c.db.Exec(query, uuid, con.Name, con.Email, con.Age, con.Address, createdOn, createdOn)
	if err != nil {
		return web.Contact{}, err
	}
	return web.Contact{
		GUID:      uuid.String(),
		Name:      con.Name,
		Address:   con.Address,
		Age:       con.Age,
		Email:     con.Email,
		CreatedOn: createdOn,
		UpdatedOn: createdOn,
	}, err
}

func (c contactRepository) Delete(id string) error {
	query := `
	DELETE FROM contacts WHERE id=$1;`
	_, err := c.db.Exec(query, id)
	return err
}

func (c contactRepository) Update(id string, con web.RequestContact) (web.Contact, error) {
	query := `
	UPDATE contacts
	SET name=$1,email=$2,age=$3,address=$4,updated_on=$5
	WHERE id=$6;`
	updatedOn := time.Now().UTC()
	_, err := c.db.Exec(query, con.Name, con.Email, con.Age, con.Address, updatedOn, id)
	return web.Contact{
		Name:      con.Name,
		Address:   con.Address,
		Age:       con.Age,
		Email:     con.Email,
		UpdatedOn: updatedOn,
	}, err
}

type contactEntity struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Age       int       `db:"age"`
	Address   string    `db:"address"`
	CreatedOn time.Time `db:"created_on"`
	UpdatedOn time.Time `db: "updated_on"`
}

func adaptToContact(entity contactEntity) web.Contact {
	return web.Contact{
		GUID:      entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Age:       entity.Age,
		Address:   entity.Address,
		CreatedOn: entity.CreatedOn,
		UpdatedOn: entity.UpdatedOn,
	}
}

func NewContactRepository(db *sql.DB) web.ContactRepository {
	return contactRepository{db: db}
}

type contactRepository struct {
	db *sql.DB
}
