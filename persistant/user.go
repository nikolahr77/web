package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"time"
)

func (u userRepository) Get(guid string) (web.User, error) {
	//getUser := `
	//SELECT * FROM users WHERE guid = $1`
	//
	return web.User{}, nil
}

func (u userRepository) Create(usr web.RequestUser) (web.User, error) {
	createUser := `
	INSERT INTO users (guid, name, password, email, created_on, age)
	VALUES ($1, $2, $3, $4, $5, $6)`
	guid := uuid.New()
	createdOn := time.Now().UTC()
	_, err := u.db.Exec(createUser, guid, usr.Name, usr.Password, usr.Email, createdOn, usr.Age)
	return web.User{
		GUID:      guid.String(),
		Name:      usr.Name,
		Password:  usr.Password,
		Email:     usr.Email,
		Age:       usr.Age,
		CreatedOn: createdOn,
	}, err
}

func (u userRepository) Update(guid string, usr web.RequestUser) (web.User, error) {
	return web.User{}, nil
}

func (u userRepository) Delete(guid string) error {
	return nil
}

type userEntity struct {
	GUID      string    `db:"guid"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	Age       int       `db:"age"`
	CreatedOn time.Time `db:"created_on"`
	Email     string    `db:"email"`
}

func NewUserRepository(db *sql.DB) web.UserRepository {
	return userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}
