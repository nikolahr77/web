package persistant_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
)

func TestCreateUserRepository(t *testing.T) {

	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	cr := persistant.NewUserRepository(DB, clock)

	newUser := web.RequestUser{
		Name:     "toni3312",
		Email:    "toncho@abv.bg",
		Age:      32,
		Password: "55f21",
	}
	actual, err := cr.Create(newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	expected := web.User{
		GUID:      actual.GUID,
		Name:      "toni3312",
		Email:     "toncho@abv.bg",
		Age:       32,
		Password:  "55f21",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
	}

	assert.Equal(t, expected, actual)
	DBCleaner(DB, "users")
}

//
//func TestUserRepository_Get(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	rows := sqlmock.NewRows([]string{"guid", "name", "password", "email", "age", "created_on", "updated_on"}).
//		AddRow("3f2-fds-d2", "Ivo", "5lJm2Sy2dkv2uxX9FcrobuZCl8WnvZ6z7yLvlt3w.ps9HZLxZv2MK", "ivo@abv.bg", 35, time.Unix(10, 0).UTC(), time.Unix(10, 0).UTC())
//
//	mock.ExpectQuery("SELECT \\* FROM users").
//		WithArgs("15").
//		WillReturnRows(rows)
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//
//	myDB := persistant.NewUserRepository(db, clock)
//
//	actual, err := myDB.Get("15")
//
//	expected := web.User{
//		GUID:      "3f2-fds-d2",
//		Name:      "Ivo",
//		Password:  "5lJm2Sy2dkv2uxX9FcrobuZCl8WnvZ6z7yLvlt3w.ps9HZLxZv2MK",
//		Email:     "ivo@abv.bg",
//		Age:       35,
//		CreatedOn: time.Unix(10, 0).UTC(),
//		UpdatedOn: time.Unix(10, 0).UTC(),
//	}
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestUserRepositoryGetReturnQueryError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectQuery("SELECT \\* FROM users").
//		WithArgs("32ff-sad2-fg5").
//		WillReturnError(SQLerror{"SQL Error"})
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//
//	myDB := persistant.NewUserRepository(db, clock)
//
//	_, err = myDB.Get("32ff-sad2-fg5")
//	expectedError := SQLerror{"SQL Error"}
//
//	assert.Equal(t, expectedError, err)
//}
//
//func TestUserRepository_Update(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	newUsr := web.RequestUser{
//		Name:     "Toni",
//		Email:    "ton@abv.bg",
//		Password: "8brbvhue9rv4h7w7845ghrvfdhz923r442",
//		Age:      73,
//	}
//
//	mock.ExpectExec("UPDATE users").WithArgs("15", newUsr).WillReturnResult(sqlmock.NewResult(0, 1))
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//
//	myDB := persistant.NewUserRepository(db, clock)
//
//	actual, err := myDB.Update("15", newUsr)
//
//	expected := web.User{
//		Name:      "Toni",
//		Email:     "ton@abv.bg",
//		Password:  "8brbvhue9rv4h7w7845ghrvfdhz923r442",
//		Age:       73,
//		UpdatedOn: time.Now().UTC(),
//	}
//
//	assert.Equal(t, expected, actual)
//}
//
//func TestUserRepository_UpdateReturnError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("UPDATE users").WillReturnError(SQLerror{"ERROR"})
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//
//	myDB := persistant.NewUserRepository(db, clock)
//
//	_, err = myDB.Update("15", web.RequestUser{})
//
//	expected := SQLerror{"ERROR"}
//
//	assert.Equal(t, expected, err)
//}
//
//func TestUserRepository_DeleteReturnError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("DELETE FROM users ").WithArgs("71241vb253fdsv").WillReturnError(SQLerror{"ERROR"})
//
//	rc := persistant.RealClock{}
//	clock := persistant.Clock(rc)
//
//	myDB := persistant.NewUserRepository(db, clock)
//
//	err = myDB.Delete("71241vb253fdsv")
//
//	expected := SQLerror{"ERROR"}
//	assert.Equal(t, expected, err)
//}

//func TestUserRepository_Delete(t *testing.T) {
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
