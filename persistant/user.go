package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"time"
)

func (u userRepository) Get(guid string) (web.User, error) {
	getUser := `
	SELECT * FROM users WHERE guid = $1`

	var ue userEntity
	rows, err := u.db.Query(getUser, guid)
	if err != nil {
		return web.User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ue.GUID, &ue.Name, &ue.Password, &ue.Email, &ue.CreatedOn, &ue.Age)
		if err != nil {
			return web.User{}, err
		}
	}
	result := adaptToUser(ue)
	return result, err

	return web.User{}, err
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
	updateUser := `
	UPDATE users
	SET name=$1, password=$2, email=$3, age=$4`
	_, err := u.db.Exec(updateUser, usr.Name, usr.Password, usr.Email, usr.Age)
	return web.User{
		Name:     usr.Name,
		Password: usr.Password,
		Email:    usr.Email,
		Age:      usr.Age,
	}, err
}

func (u userRepository) Delete(guid string) error {
	deleteUser := `
	DELETE FROM users WHERE guid = $1`

	_, err := u.db.Exec(deleteUser, guid)
	return err
}

type userEntity struct {
	GUID      string    `db:"guid"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	Age       int       `db:"age"`
	CreatedOn time.Time `db:"created_on"`
	Email     string    `db:"email"`
}

func adaptToUser(u userEntity) web.User {
	return web.User{
		GUID:      u.GUID,
		Name:      u.Name,
		Password:  u.Password,
		Email:     u.Email,
		Age:       u.Age,
		CreatedOn: u.CreatedOn,
	}
}

func NewUserRepository(db *sql.DB) web.UserRepository {
	return userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}
