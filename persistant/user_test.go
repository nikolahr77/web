package persistant_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
	"time"
)

func TestUserRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"guid", "name", "password", "email", "age", "created_on", "updated_on"}).
		AddRow("3f2-fds-d2", "Ivo", "5lJm2Sy2dkv2uxX9FcrobuZCl8WnvZ6z7yLvlt3w.ps9HZLxZv2MK", "ivo@abv.bg", 35, time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())

	mock.ExpectQuery("SELECT \\* FROM users").
		WithArgs("15").
		WillReturnRows(rows)

	myDB := persistant.NewUserRepository(db)

	actual, err := myDB.Get("15")

	expected := web.User{
		GUID:      "3f2-fds-d2",
		Name:      "Ivo",
		Password:  "5lJm2Sy2dkv2uxX9FcrobuZCl8WnvZ6z7yLvlt3w.ps9HZLxZv2MK",
		Email:     "ivo@abv.bg",
		Age:       35,
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(10, 0).UTC(),
	}

	assert.Equal(t, expected, actual)
}

func TestUserRepositoryGetReturnQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM users").
		WithArgs("32ff-sad2-fg5").
		WillReturnError(SQLerror{"SQL Error"})

	myDB := persistant.NewUserRepository(db)

	_, err = myDB.Get("32ff-sad2-fg5")
	expectedError := SQLerror{"SQL Error"}

	assert.Equal(t, expectedError, err)
}
