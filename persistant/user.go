package persistant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/web"
	"github.com/web/convert"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//GetByName searches the DB for a user with the specified name
// and returns this user
func (u userRepository) GetByName(name string) (web.User, error) {
	query := `
		SELECT * FROM users WHERE name = $1`

	rows, err := u.db.Query(query, name)

	var ue userEntity

	if err != nil {
		return web.User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ue.GUID, &ue.Name, &ue.Password, &ue.Email, &ue.Age, &ue.CreatedOn, &ue.UpdatedOn)
		if err != nil {
			return web.User{}, err
		}
	}
	result := web.User{}
	convert.SourceToDestination(ue, &result)
	return result, err
}

//Get is used to return a user from the DB by a given ID.
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
		err := rows.Scan(&ue.GUID, &ue.Name, &ue.Password, &ue.Email, &ue.Age, &ue.CreatedOn, &ue.UpdatedOn)
		if err != nil {
			return web.User{}, err
		}
	}
	result := web.User{}
	convert.SourceToDestination(ue, &result)
	return result, err

	return web.User{}, err
}

//Create adds a new user to the DB
func (u userRepository) Create(usr web.RequestUser) (web.User, error) {
	createUser := `
	INSERT INTO users (guid, name, password, email, created_on, age, updated_on)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	guid := uuid.New()
	createdOn := time.Now().UTC()

	passBytes := []byte(usr.Password)
	hashedBytes, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		return web.User{}, err
	}
	cryptPass := string(hashedBytes[:])

	_, err = u.db.Exec(createUser, guid, usr.Name, cryptPass, usr.Email, createdOn, usr.Age, createdOn)
	return web.User{
		GUID:      guid.String(),
		Name:      usr.Name,
		Password:  usr.Password,
		Email:     usr.Email,
		Age:       usr.Age,
		CreatedOn: createdOn,
		UpdatedOn: createdOn,
	}, err
}

//Update searches the DB for a user by a given
// ID and updates the users with the given RequestUser
func (u userRepository) Update(guid string, usr web.RequestUser) (web.User, error) {
	updateUser := `
	UPDATE users
	SET name=$1, password=$2, email=$3, age=$4, updated_on=$5
	WHERE guid = $6`
	updatedOn := time.Now().UTC()
	saltedBytes := []byte(usr.Password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return web.User{}, err
	}
	cryptPass := string(hashedBytes[:])

	_, err = u.db.Exec(updateUser, usr.Name, cryptPass, usr.Email, usr.Age, updatedOn, guid)
	return web.User{
		Name:      usr.Name,
		Password:  usr.Password,
		Email:     usr.Email,
		Age:       usr.Age,
		UpdatedOn: updatedOn,
	}, err
}

//Delete is used to remove a user from the DB by a given ID.
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
	UpdatedOn time.Time `db:"updated_on"`
	Email     string    `db:"email"`
}

func NewUserRepository(db *sql.DB) web.UserRepository {
	return userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}
