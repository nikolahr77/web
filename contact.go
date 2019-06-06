package web

type Contact struct {
	ID		int64
	Name    string
	Email   string
	Age     int
	Address string
}

type ContactRepository interface {
	Get(id int64) (Contact, error)
	Create(con Contact) (Contact, error)
	Delete(id int64) error
	Update(id int64, con Contact) (Contact, error)
}
