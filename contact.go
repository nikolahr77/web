package web

import "time"

type Contact struct {
	GUID    string
	Name    string
	Email   string
	Age     int
	Address string
	CreatedOn time.Time
	UpdatedOn time.Time
}

type ContactRepository interface {
	//Get(id string) (Contact, error)
	Create(con Contact) (Contact, error)
	//Delete(id string) error
	//Update(id string, con Contact) (Contact, error)
}
