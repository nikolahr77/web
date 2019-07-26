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

var DB *sql.DB
var Cleaner = dbcleaner.New()

func DBCleaner(db *sql.DB, table string) {
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

func TestCreateContactRepository(t *testing.T) {

	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	cr := persistant.NewContactRepository(DB, clock)

	newContact := web.RequestContact{
		Name:    "Dani",
		Email:   "dani@abv.bg",
		Age:     62,
		Address: "Pleven",
	}
	userID := uuid.New()
	actual, err := cr.Create(newContact, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Contact{
		GUID:      actual.GUID,
		Name:      "Dani",
		Email:     "dani@abv.bg",
		Age:       62,
		Address:   "Pleven",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
		UserID:    userID.String(),
	}

	assert.Equal(t, expected, actual)

	DBCleaner(DB, "contacts")
}

func TestUpdateContactRepository(t *testing.T) {

	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

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

	actual, err := cr.Update(contactToUpdate.GUID, newContact, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Contact{
		GUID:      actual.GUID,
		Name:      "Ivan",
		Email:     "ivo@abv.bg",
		Age:       32,
		Address:   "Yambol",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
		UserID:    actual.UserID,
	}

	assert.Equal(t, expected, actual)

	DBCleaner(DB, "contacts")
}

func TestDeleteContactRepository(t *testing.T) {

	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

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

	assert.Equal(t, err, nil)
}

func TestGetContactRepository(t *testing.T) {
	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	cr := persistant.NewContactRepository(DB, clock)

	oldContact := web.RequestContact{
		Name:    "Dani",
		Email:   "dani@abv.bg",
		Age:     62,
		Address: "Pleven",
	}

	userID := uuid.New()
	contactToGet, err := cr.Create(oldContact, userID.String())
	if err != nil {
		panic(err)
	}

	actual, err := cr.Get(contactToGet.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Contact{
		GUID:      actual.GUID,
		Name:      "Dani",
		Email:     "dani@abv.bg",
		Age:       62,
		Address:   "Pleven",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
		UserID:    userID.String(),
	}

	assert.Equal(t, expected, actual)
	DBCleaner(DB, "contacts")
}

//
//func TestContactRepository_Get(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	rows := sqlmock.NewRows([]string{"id", "name", "email", "age", "address", "created_on", "updated_on"}).
//		AddRow("15", "Ivan", "ivan@abv.bg", 15, "Sofia", time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())
//
//	mock.ExpectQuery("SELECT \\* FROM contacts").
//		WithArgs("15").
//		WillReturnRows(rows)
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//	myDB := persistant.NewContactRepository(db, clock)
//
//	actual, err := myDB.Get("15", )
//
//	expected := web.Contact{
//		GUID:      "15",
//		Name:      "Ivan",
//		Email:     "ivan@abv.bg",
//		Age:       15,
//		Address:   "Sofia",
//		CreatedOn: time.Unix(10, 0).UTC(),
//		UpdatedOn: time.Unix(10, 0).UTC(),
//	}
//}
//
//func ContactRepositoryTest(db *sql.DB, clock persistant.Clock) web.ContactRepository {
//	return contactRepository{db: db, clock: clock}
//}
//
//type contactRepository struct {
//	db    *sql.DB
//	clock persistant.Clock
//}

//assert.Equal(t, expected, actual)
//}
//
//func TestContactRepositoryGetReturnQueryError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectQuery("SELECT \\* FROM contacts").
//		WithArgs("15").
//		WillReturnError(SQLerror{"SQL Error"})
//
//	myDB := persistant.NewContactRepository(db)
//
//	_, err = myDB.Get("15")
//	expectedError := SQLerror{"SQL Error"}
//
//	assert.Equal(t, expectedError, err)
//
//}

//func TestContactRepository_GetReturnRowsError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectQuery("SELECT \\* FROM contacts").
//		WithArgs("15").
//		WillReturnRows(nil)
//
//	myDB := persistant.NewContactRepository(db)
//
//	actual, err := myDB.Get("15")
//
//	//expected := errors.New("SQL")
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestContactRepository_Update(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	newContact := web.RequestContact{
//		Name:    "Petur",
//		Email:   "petur@abv.bg",
//		Age:     95,
//		Address: "Plovdiv",
//	}
//
//	mock.ExpectExec("UPDATE contacts").WithArgs("15", newContact).WillReturnResult(sqlmock.NewResult(0, 1))
//
//	myDB := persistant.NewContactRepository(db)
//
//	actual, err := myDB.Update("15", newContact)
//
//	expected := web.Contact{
//		Name:      "Petur",
//		Email:     "petur@abv.bg",
//		Age:       95,
//		Address:   "Plovdiv",
//		UpdatedOn: time.Unix(10, 0).UTC(),
//	}
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestContactRepository_UpdateReturnError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("UPDATE contacts").WillReturnError(SQLerror{"ERROR"})
//
//	myDB := persistant.NewContactRepository(db)
//
//	_, err = myDB.Update("15", web.RequestContact{})
//
//	expected := SQLerror{"ERROR"}
//
//	assert.Equal(t, expected, err)
//}
//
//func TestContactRepository_Create(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	newContact := web.RequestContact{
//		Name:    "Dani",
//		Email:   "dani@abv.bg",
//		Age:     62,
//		Address: "Pleven",
//	}
//
//	mock.ExpectExec("INSERT INTO contacts").WithArgs(newContact).WillReturnResult(sqlmock.NewResult(0, 1))
//
//	myDB := persistant.NewContactRepository(db)
//
//	actual, err := myDB.Create(newContact)
//
//	expected := web.Contact{
//		GUID:      "32b-15",
//		Name:      "Dani",
//		Email:     "dani@abv.bg",
//		Age:       62,
//		Address:   "Pleven",
//		CreatedOn: time.Now(),
//		UpdatedOn: time.Now(),
//	}
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestContactRepository_CreateReturnError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("INSERT INTO contacts").WillReturnError(SQLerror{"ERROR"})
//	mock.ExpectQuery("SELECT \\* FROM contacts").
//		WithArgs("15").
//		WillReturnError(SQLerror{"ERROR"})
//
//	myDB := persistant.NewContactRepository(db)
//
//	_, err = myDB.Create(web.RequestContact{})
//
//	expected := SQLerror{"ERROR"}
//
//	assert.Equal(t, expected, err)
//}
//
//func TestContactRepository_Delete(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("DELETE FROM contacts ").WithArgs("52-23d").WillReturnResult(sqlmock.NewResult(0, 1))
//
//	myDB := persistant.NewContactRepository(db)
//
//	err = myDB.Delete("52-23d")
//
//	assert.Equal(t, nil, err)
//}
//
//func TestContactRepository_DeleteReturnError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("DELETE FROM contacts ").WithArgs("52-23d").WillReturnError(SQLerror{"ERROR"})
//
//	myDB := persistant.NewContactRepository(db)
//
//	err = myDB.Delete("52-23d")
//	expected := SQLerror{"ERROR"}
//
//	assert.Equal(t, expected, err)
//}
