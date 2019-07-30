package persistant_test

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"gopkg.in/khaiql/dbcleaner.v2"
	"log"
	"os"
	"testing"
	"time"
)

//import (
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/stretchr/testify/assert"
//	"github.com/web"
//	"github.com/web/persistant"
//	"testing"
//	"time"
//)

type SQLerror struct {
	content string
}

func (err SQLerror) Error() string {
	return "SQL Error"
}

type fakeClock struct {
	Seconds int64
}

func (f fakeClock) Now() time.Time {
	return time.Unix(f.Seconds, 0).UTC()
}

var DB *sql.DB
var Cleaner = dbcleaner.New()

func dbCleaner(db *sql.DB, table string) {
	query := fmt.Sprintf(`DELETE FROM %s`, table)
	_, err := db.Exec(query)
	if err != nil {
		log.Print(err)
	}
}

func TestMain(m *testing.M) {
	connStr := "user=postgres dbname=testmail sslmode=disable password=1234"
	DB, _ = sql.Open("postgres", connStr)
	defer DB.Close()
	code := m.Run()
	os.Exit(code)
}

func TestCreateUpdateDeleteContact(t *testing.T) {
	dbCleaner(DB, "contacts")

	clock := fakeClock{
		Seconds: 25000,
	}

	cr := persistant.NewContactRepository(DB, clock)

	oldContact := web.RequestContact{
		Name:    "Dani",
		Email:   "dani@abv.bg",
		Age:     62,
		Address: "Pleven",
	}

	newContact := web.RequestContact{
		Name:    "Ivan",
		Email:   "ivo@abv.bg",
		Age:     32,
		Address: "Yambol",
	}
	userID := uuid.New()
	contactToUpdate, err := cr.Create(oldContact, userID.String())
	if err != nil {
		panic(err)
	}

	_, err = cr.Update(contactToUpdate.GUID, newContact, userID.String())
	if err != nil {
		panic(err)
	}

	actual, err := cr.Get(contactToUpdate.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Contact{
		GUID:      actual.GUID,
		Name:      "Ivan",
		Email:     "ivo@abv.bg",
		Age:       32,
		Address:   "Yambol",
		CreatedOn: time.Unix(25000, 0).UTC(),
		UpdatedOn: time.Unix(25000, 0).UTC(),
		UserID:    actual.UserID,
	}

	assert.Equal(t, expected, actual)
}

func TestCreateDeleteGetContact(t *testing.T) {
	dbCleaner(DB, "contacts")

	clock := fakeClock{
		Seconds: 25000,
	}

	cr := persistant.NewContactRepository(DB, clock)

	oldContact := web.RequestContact{
		Name:    "Dani",
		Email:   "dani@abv.bg",
		Age:     62,
		Address: "Pleven",
	}

	userID := uuid.New()
	old, err := cr.Create(oldContact, userID.String())
	if err != nil {
		panic(err)
	}

	err = cr.Delete(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	actual, err := cr.Get(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, web.Contact{})
}

func TestCreateContactWithSameEmail(t *testing.T) {
	dbCleaner(DB, "contacts")

	clock := fakeClock{
		Seconds: 25000,
	}

	cr := persistant.NewContactRepository(DB, clock)

	Contact := web.RequestContact{
		Name:    "Dani",
		Email:   "ivo@abv.bg",
		Age:     62,
		Address: "Pleven",
	}

	newContact := web.RequestContact{
		Name:    "Ivailo",
		Email:   "ivo@abv.bg",
		Age:     32,
		Address: "Yambol",
	}
	userID := uuid.New()

	_, err := cr.Create(Contact, userID.String())
	if err != nil {
		panic(err)
	}

	_, err = cr.Create(newContact, userID.String())

	assert.Contains(t, err.Error(), "duplicate key value")
}
