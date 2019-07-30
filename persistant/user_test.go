package persistant_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestCreateUpdateGetUserRepository(t *testing.T) {
	dbCleaner(DB, "users")

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
	_, err = cr.Update(old.GUID, newUser)
	if err != nil {
		panic(err)
	}

	actual, err := cr.Get(old.GUID)
	if err != nil {
		panic(err)
	}

	password := []byte(actual.Password)
	requestPassword := []byte("5ggfdsh52f221")
	err = bcrypt.CompareHashAndPassword(password, requestPassword)
	if err != nil {
		panic(err)
	}

	expected := web.User{
		GUID:      actual.GUID,
		Name:      "misho55",
		Email:     "mishoo@abv.bg",
		Age:       12,
		Password:  actual.Password,
		CreatedOn: time.Unix(25000, 0).UTC(),
		UpdatedOn: time.Unix(25000, 0).UTC(),
	}

	assert.Equal(t, expected, actual)
	assert.Equal(t, err, nil)
}

func TestCreateDeleteGetUserRepository(t *testing.T) {
	dbCleaner(DB, "users")

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

	actual, err := cr.Get(old.GUID)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, web.User{})
}
