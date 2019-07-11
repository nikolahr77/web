package persistant_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
	"time"
)

func TestContactRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "age", "address", "created_on", "updated_on"}).
		AddRow("15", "Ivan", "ivan@abv.bg", 15, "Sofia", time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())

	mock.ExpectQuery("SELECT \\* FROM contacts").
		WithArgs("15").
		WillReturnRows(rows)

	myDB := persistant.NewContactRepository(db)

	actual, err := myDB.Get("15")

	expected := web.Contact{
		GUID:      "15",
		Name:      "Ivan",
		Email:     "ivan@abv.bg",
		Age:       15,
		Address:   "Sofia",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(10, 0).UTC(),
	}

	assert.Equal(t, expected, actual)
}

func TestContactRepositoryGetReturnQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM contacts").
		WithArgs("15").
		WillReturnError(errors.New("test error"))

	myDB := persistant.NewContactRepository(db)

	actual, err := myDB.Get("15")

	expected := web.Contact{}
	assert.Equal(t, expected, actual)
}
