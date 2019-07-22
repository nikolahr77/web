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
	UserID    string
}

type RequestContact struct {
	Name    string
	Email   string
	Age     int
	Address string
	UserID  string
}

type ContactRepository interface {
	GetAll(camSegmentation Segmentation) ([]Contact, error)
	Get(id string) (Contact, error)
	Create(con RequestContact, userID string) (Contact, error)
	Delete(id string) error
	Update(id string, con RequestContact) (Contact, error)
}
