package web

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	UserRepository UserRepository
}

func (am AuthMiddleware) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/users" {
			next.ServeHTTP(w, r)
			return
		}
		usr, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "authentication failed", http.StatusUnauthorized)
			return
		}
		user, err := am.UserRepository.GetByName(usr)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		passByte := []byte(user.Password)
		requestPassByte := []byte(pass)
		err = bcrypt.CompareHashAndPassword(passByte, requestPassByte)
		if err != nil {
			log.Println(err)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type MessageRequest struct {
	Messages []Message2
}

type Message2 struct {
	From     email   //az
	To       []email //contacts
	TextPart string  //slagame content ot messages
}

type email struct {
	email string
	name  string
}

//func GetUserByName(name string) (User, error) {
//	return User{}, nil
//}
//func Get(guid string) (User, error) {
//	return User{}, nil
//}
//func Create(usr RequestUser) (User, error) {
//	return User{}, nil
//}
//func Update(guid string, usr RequestUser) (User, error) {
//	return User{}, nil
//}
//func Delete(guid string) error {
//	return nil
//}
//
//func NewUserRepository(db *sql.DB) UserRepository {
//	return userRepository{db: db}
//}
//type userRepository struct {
//	db *sql.DB
//}

//func validate(username, password string) bool {
//	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//		return false
//	}
//
//	searchName := `
//	SELECT password FROM users WHERE name = $1`
//
//	usr, err := db.Query(searchName, username)
//
//	var u user
//
//	defer usr.Close()
//	for usr.Next() {
//		err := usr.Scan(&u.Password)
//		if err != nil {
//			log.Println(err)
//			return false
//		}
//	}
//	fmt.Println(u.Password)
//	fmt.Println(password)
//
//	passByte := []byte(u.Password)
//	requestPassByte := []byte(password)
//	err = bcrypt.CompareHashAndPassword(passByte, requestPassByte)
//	if err != nil {
//		log.Println(err)
//		return false
//	}
//	return true
//}
//
