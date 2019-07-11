package persistant_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
	"time"
)

func TestMessageRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"guid", "name", "content", "created_on", "updated_on"}).
		AddRow("3ff-d2", "Welcome", "This is a welcome msg", time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())

	mock.ExpectQuery("SELECT \\* FROM messages").
		WithArgs("15").
		WillReturnRows(rows)

	myDB := persistant.NewMessageRepository(db)

	actual, err := myDB.Get("15")

	expected := web.Message{
		GUID:      "3ff-d2",
		Name:      "Welcome",
		Content:   "This is a welcome msg",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(10, 0).UTC(),
	}

	assert.Equal(t, expected, actual)
}

func TestMessageRepositoryGetReturnQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM messages").
		WithArgs("15").
		WillReturnError(SQLerror{"SQL Error"})

	myDB := persistant.NewMessageRepository(db)

	_, err = myDB.Get("15")
	expectedError := SQLerror{"SQL Error"}

	assert.Equal(t, expectedError, err)

}

func TestMessageRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"guid", "name", "content", "created_on", "updated_on"}).
		AddRow("3ff-d2", "Welcome", "This is a welcome msg", time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())

	newMsg := web.RequestMessage{
		Name:    "Edited",
		Content: "This message was edited",
	}

	mock.ExpectExec("UPDATE messages").WithArgs("15", newMsg).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectQuery("SELECT \\* FROM messages").
		WithArgs("15").
		WillReturnRows(rows)

	myDB := persistant.NewMessageRepository(db)

	actual, err := myDB.Update("15", newMsg)

	expected := web.Message{
		Name:      "Edited",
		Content:   "This message was edited",
		UpdatedOn: time.Now().UTC(),
	}

	assert.Equal(t, expected, actual)
}

func TestMessageRepository_UpdateReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE messages").WillReturnError(SQLerror{"ERROR"})
	mock.ExpectQuery("SELECT \\* FROM messages").
		WithArgs("15").
		WillReturnError(SQLerror{"ERROR"})

	myDB := persistant.NewMessageRepository(db)

	_, err = myDB.Update("15", web.RequestMessage{})

	expected := SQLerror{"ERROR"}

	assert.Equal(t, expected, err)
}
