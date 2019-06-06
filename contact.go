package web

type Contact struct {
	Name    string
	Email   string
	Age     int
	Address string
}

type ContactRepository interface {
	//Get(id int64) (*sql.Rows,error)
	Create(con Contact) (Contact, error)
	Delete(id int64) error
	//Update(contact Contact) (Contact,error)
}
