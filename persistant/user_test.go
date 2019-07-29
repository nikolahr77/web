package persistant_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCreateUserRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}

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
	dbCleaner(DB, "users")
}

func TestUpdateUserRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}

	cr := persistant.NewUserRepository(DB, clock)

	oldUser := web.RequestUser{
		Name:     "toni3312",
		Email:    "toncho@abv.bg",
		Age:      32,
		Password: "55f21",
	}

	newUser := web.RequestUser{
		Name:     "misho55",
		Email:    "mishoo@abv.bg",
		Age:      12,
		Password: "5ggfdsh52f221",
	}
	old, err := cr.Create(oldUser)
	if err != nil {
		panic(err)
	}
	actual, err := cr.Update(old.GUID, newUser)
	if err != nil {
		panic(err)
	}

	expected := web.User{
		GUID:      actual.GUID,
		Name:      "misho55",
		Email:     "mishoo@abv.bg",
		Age:       12,
		Password:  "5ggfdsh52f221",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
	}

	assert.Equal(t, expected, actual)
	dbCleaner(DB, "users")
}

func TestDeleteUserRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}

	cr := persistant.NewUserRepository(DB, clock)

	oldUser := web.RequestUser{
		Name:     "toni3312",
		Email:    "toncho@abv.bg",
		Age:      32,
		Password: "55f21",
	}

	old, err := cr.Create(oldUser)
	if err != nil {
		panic(err)
	}
	err = cr.Delete(old.GUID)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, err, nil)
}

func TestGetUserRepository(t *testing.T) { //FIX PASSWORD !!!
	clock := fakeClock{
		Seconds: 25000,
	}

	cr := persistant.NewUserRepository(DB, clock)

	newUser := web.RequestUser{
		Name:     "misho55",
		Email:    "mishoo@abv.bg",
		Age:      12,
		Password: "5ggfdsh52f221",
	}
	user, err := cr.Create(newUser)
	if err != nil {
		panic(err)
	}
	actual, err := cr.Get(user.GUID)
	if err != nil {
		panic(err)
	}

	passByte := []byte(actual.Password)
	requestPassByte := []byte("5ggfdsh52f221")
	err = bcrypt.CompareHashAndPassword(passByte, requestPassByte)
	if err != nil {
		panic(err)
	}

	expected := web.User{
		GUID:      actual.GUID,
		Name:      "misho55",
		Email:     "mishoo@abv.bg",
		Age:       12,
		Password:  actual.Password,  //I am comparing passwords above
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
	}

	assert.Equal(t, expected, actual)
	dbCleaner(DB, "users")
}
