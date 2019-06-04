package web

type Contact struct {
	Name    string
	Email   string
	Age     int
	Address string
}

type ContactRepository interface {
	Get(id int64) (Contact,error)
	Post(id int64) (Contact,error)
	Delete(id int64) (Contact,error)
}