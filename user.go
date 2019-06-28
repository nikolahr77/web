package web

import "time"

type User struct {
	GUID      string
	Name      string
	Password  string
	Email     string
	Age       int
	CreatedOn time.Time
	UpdatedOn time.Time
}

type RequestUser struct {
	Name     string
	Email    string
	Age      int
	Password string
}

type UserRepository interface {
	Get(guid string) (User, error)
	Create(usr RequestUser) (User, error)
	Update(guid string, usr RequestUser) (User, error)
	Delete(guid string) error
}
