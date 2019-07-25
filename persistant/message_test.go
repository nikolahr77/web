package persistant_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"log"
	"testing"
)

func TestCreateMessageRepository(t *testing.T) {

	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	mr := persistant.NewMessageRepository(DB, clock)

	newMsg := web.RequestMessage{
		Name:    "TestMSG",
		Content: "This is a test message",
	}
	userID := uuid.New()
	actual, err := mr.Create(newMsg, userID.String())
	if err != nil {
		log.Print(err)
	}
	fmt.Println(err)
	expected := web.Message{
		GUID:      actual.GUID,
		Name:      "TestMSG",
		Content:   "This is a test message",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
		UserID:    userID.String(),
	}

	assert.Equal(t, expected, actual)
}

//func TestMessageRepository_Get(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	rows := sqlmock.NewRows([]string{"guid", "name", "content", "created_on", "updated_on"}).
//		AddRow("3ff-d2", "Welcome", "This is a welcome msg", time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())
//
//	mock.ExpectQuery("SELECT \\* FROM messages").
//		WithArgs("15").
//		WillReturnRows(rows)
//
//	myDB := persistant.NewMessageRepository(db)
//
//	actual, err := myDB.Get("15")
//
//	expected := web.Message{
//		GUID:      "3ff-d2",
//		Name:      "Welcome",
//		Content:   "This is a welcome msg",
//		CreatedOn: time.Unix(10, 0).UTC(),
//		UpdatedOn: time.Unix(10, 0).UTC(),
//	}
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestMessageRepositoryGetReturnQueryError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectQuery("SELECT \\* FROM messages").
//		WithArgs("15").
//		WillReturnError(SQLerror{"SQL Error"})
//
//	myDB := persistant.NewMessageRepository(db)
//
//	_, err = myDB.Get("15")
//	expectedError := SQLerror{"SQL Error"}
//
//	assert.Equal(t, expectedError, err)
//
//}
//
//func init() {
//	// Uncomment and add to _test.go init()
//
//}
//
//var timeNow = time.Now
//
//func main() {
//	fmt.Println(timeNow())
//}
//
//func init() {
//	// Uncomment and add to _test.go init()
//	timeNow = func() time.Time {
//		t, _ := time.Parse("2006-01-02 15:04:05", "2017-01-20 01:02:03")
//		return t
//	}
//}
//
//func TestMessageRepository_Update(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	newMsg := web.RequestMessage{
//		Name:    "Edited",
//		Content: "This message was edited",
//	}
//
//	mock.ExpectExec("UPDATE messages").WithArgs("15", newMsg).WillReturnResult(sqlmock.NewResult(0, 1))
//
//	myDB := persistant.NewMessageRepository(db)
//
//	actual, err := myDB.Update("15", newMsg)
//
//	expected := web.Message{
//		Name:      "Edited",
//		Content:   "This message was edited",
//		UpdatedOn: timeNow(),
//	}
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestMessageRepository_UpdateReturnError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("UPDATE messages").WillReturnError(SQLerror{"ERROR"})
//
//	myDB := persistant.NewMessageRepository(db)
//
//	_, err = myDB.Update("15", web.RequestMessage{})
//
//	expected := SQLerror{"ERROR"}
//
//	assert.Equal(t, expected, err)
//}
