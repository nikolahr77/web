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

func NewContactRepository(db *sql.DB) web.ContactRepository  {
	return contactRepository{db:db}
}

func (c contactRepository) Get(id int64) (*sql.Rows,error) {
	query := `SELECT name, email, age, address FROM contacts;`
	var e contactEntity
	rows,err := c.db.Query(query)
	if err!= nil {
		panic(err)
	}
	return rows,err
}

func adaptToContact(entity contactEntity)  web.Contact{
	return web.Contact{
		Name:entity.Name,
		Email:entity.Email,
		Age:entity.Age,
		Address:entity.Address,
	}
}

//func (c contactRepository) Create(contact web.Contact) (web.Contact,error)  {
//	query := `
//	INSERT INTO contacts (name,email,age,address)
//	VALUES ($1, $2, $3, $4);`
//	var e contactEntity
//	c.db.Exec(query,)
//}
//
//func (c contactRepository) Delete(id int64) (error)  {
//	query := `
//	DELETE FROM contacts WHERE name=$1;`
//	var e contactEntity
//	c.db.Exec(query,e)
//}
//
//func (c contactRepository)  Update(id int64) (web.Contact,error){
//	query := `
//	UPDATE contacts
//	SET name=$1,email=$2,age=$3,address=$4
//	WHERE id=$5;`
//	var e contactEntity
//	c.db.Exec(query,e)
//}



