package persistant

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/web"
	"github.com/web/convert"
	"time"
)

//GetAll returns all of the contacts which have the wanted address and age
func (c contactRepository) GetAll(camSegmentation web.Segmentation) ([]web.Contact, error) {
	query := `
	SELECT name, email FROM contacts WHERE age = $1 AND address = $2`
	var e contactEntity
	rows, err := c.db.Query(query, camSegmentation.Age, camSegmentation.Address)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var contacts []contactEntity
	for rows.Next() {
		err := rows.Scan(&e.Name, &e.Email)
		if err != nil {
			return []web.Contact{}, err
		}
		contacts = append(contacts, e)
	}
	result := []web.Contact{}
	convert.SourceToDestination(contacts, &result)
	return result, err
}

//Get is used to return a contact from the DB by a given ID.
func (c contactRepository) Get(id string, userID string) (web.Contact, error) {
	query := `SELECT * FROM contacts WHERE id=$1 AND userid = $2`
	var e contactEntity
	rows, err := c.db.Query(query, id, userID)
	if err != nil {
		return web.Contact{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.GUID, &e.Name, &e.Email, &e.Age, &e.Address, &e.CreatedOn, &e.UpdatedOn, &e.UserID)
		if err != nil {
			return web.Contact{}, err
		}
	}
	result := web.Contact{}
	convert.SourceToDestination(e, &result)
	return result, err
}

//Create adds a new contact to the DB
func (c contactRepository) Create(con web.RequestContact, userID string) (web.Contact, error) {
	uuid := uuid.New()
	query := `
	INSERT INTO contacts (id, name,email,age,address,created_on,updated_on, userID)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	createdOn := c.clock.Now().UTC()
	_, err := c.db.Exec(query, uuid, con.Name, con.Email, con.Age, con.Address, createdOn, createdOn, userID)
	if err != nil {
		fmt.Println(err)
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
		UserID:    userID,
	}, err
}

//Delete is used to remove a contact from the DB by a given ID.
func (c contactRepository) Delete(id string, userID string) error {
	query := `
	DELETE FROM contacts WHERE id=$1 AND userID = $2;`
	_, err := c.db.Exec(query, id, userID)
	return err
}

//Update searches the DB for a contact by a given
// ID and updates the campaign with the given RequestContact
func (c contactRepository) Update(id string, con web.RequestContact, userID string) (web.Contact, error) {
	query := `
	UPDATE contacts
	SET name=$1,email=$2,age=$3,address=$4,updated_on=$5
	WHERE id=$6 AND userID = $7;` //da dobavq i userID
	updatedOn := c.clock.Now().UTC()
	_, err := c.db.Exec(query, con.Name, con.Email, con.Age, con.Address, updatedOn, id, userID)
	return web.Contact{
		Name:      con.Name,
		Address:   con.Address,
		Age:       con.Age,
		Email:     con.Email,
		UpdatedOn: updatedOn,
	}, err
}

type contactEntity struct {
	GUID      string    `db:"uuid"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Age       int       `db:"age"`
	Address   string    `db:"address"`
	CreatedOn time.Time `db:"created_on"`
	UpdatedOn time.Time `db:"updated_on"`
	UserID    string    `db: "userID"`
}

func NewContactRepository(db *sql.DB, clock Clock) web.ContactRepository {
	return contactRepository{db: db, clock: clock}
}

type contactRepository struct {
	db    *sql.DB
	clock Clock
}
