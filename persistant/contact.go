package persistant

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/web"
	"time"
)

type contactRepository struct {
	db *sql.DB
}

type contactEntity struct {
	ID      string  `db:"id"`
	Name    string `db:"name"`
	Email   string `db:"email"`
	Age     int    `db:"age"`
	Address string `db:"address"`
	CreatedOn time.Time `db:"created_on"`
	UpdatedOn time.Time `db: "updated_on"`
}

func NewContactRepository(db *sql.DB) web.ContactRepository {
	return contactRepository{db: db}
}
//
//func (c contactRepository) Get(id string) (web.Contact, error) {
//	query := `SELECT * FROM contacts WHERE id=$1 and user_id = 2;`
//	var e contactEntity
//	rows, err := c.db.Query(query, id)
//	if err != nil {
//		return web.Contact{}, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Address)
//		if err != nil {
//			return web.Contact{}, err
//		}
//	}
//	result := adaptToContact(e)
//	return result, err
//}

func adaptToContact(entity contactEntity) web.Contact {
	return web.Contact{
		GUID:    entity.ID,
		Name:    entity.Name,
		Email:   entity.Email,
		Age:     entity.Age,
		Address: entity.Address,
	}
}

func (c contactRepository) Create(con web.Contact) (web.Contact, error) {
	uuid := uuid.New()
	query := `
	INSERT INTO contacts (id, name,email,age,address,created_on,updated_on)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`
	createTime := time.Now
	_, err := c.db.Exec(query, uuid, con.Name, con.Email, con.Age, con.Address,createTime(),createTime())
	fmt.Println(err)
	return web.Contact{
		//GUID: uuid,
		Name: con.Name,
		Address: con. Address,
		Age: con.Age,
		Email: con.Email,
		CreatedOn: con.CreatedOn,
		UpdatedOn: con.UpdatedOn,
	}, err
}

//func (c contactRepository) Delete(id string) error {
//	query := `
//	DELETE FROM contacts WHERE id=$1;`
//	//var e contactEntity
//	_, err := c.db.Exec(query, id)
//	return err
//}
//
//func (c contactRepository) Update(id string, con web.Contact) (web.Contact, error) {
//	query := `
//	UPDATE contacts
//	SET name=$1,email=$2,age=$3,address=$4
//	WHERE id=$5;`
//	_, err := c.db.Exec(query, con.Name, con.Email, con.Age, con.Address, id)
//	return con, err
//}
