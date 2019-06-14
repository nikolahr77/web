package web

import "time"

type Contact struct {
	GUID      string
	Name      string
	Email     string
	Age       int
	Address   string
	CreatedOn time.Time
	UpdatedOn time.Time
}

type RequestContact struct {
	Name      string
	Email     string
	Age       int
	Address   string
}

type ContactRepository interface {
	Get(id string) (Contact, error)
	Create(con RequestContact) (Contact, error)
	Delete(id string) error
	Update(id string, con RequestContact) (Contact, error)
}
