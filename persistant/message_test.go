package persistant_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
	"time"
)

func TestCreateUpdateGetMessageRepository(t *testing.T) {
	dbCleaner(DB, "messages")

	clock := fakeClock{
		Seconds: 25000,
	}

	mr := persistant.NewMessageRepository(DB, clock)

	oldMsg := web.RequestMessage{
		Name:    "TestMSG",
		Content: "This is a test message",
	}

	newMsg := web.RequestMessage{
		Name:    "NewMSG",
		Content: "This is the new test message",
	}
	userID := uuid.New()
	old, err := mr.Create(oldMsg, userID.String())
	if err != nil {
		panic(err)
	}
	_, err = mr.Update(old.GUID, newMsg, userID.String())
	if err != nil {
		panic(err)
	}

	actual, err := mr.Get(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Message{
		GUID:      actual.GUID,
		Name:      "NewMSG",
		Content:   "This is the new test message",
		CreatedOn: time.Unix(25000, 0).UTC(), //I should't do this
		UpdatedOn: time.Unix(25000, 0).UTC(),
		UserID:    userID.String(),
	}

	assert.Equal(t, expected, actual)
}

func TestCreateDeleteGetMessageRepository(t *testing.T) {
	dbCleaner(DB, "messages")

	clock := fakeClock{
		Seconds: 25000,
	}

	mr := persistant.NewMessageRepository(DB, clock)

	oldMsg := web.RequestMessage{
		Name:    "TestMSG",
		Content: "This is a test message",
	}

	userID := uuid.New()
	old, err := mr.Create(oldMsg, userID.String())
	if err != nil {
		panic(err)
	}
	err = mr.Delete(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	actual, err := mr.Get(old.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	assert.Equal(t, err, nil)
	assert.Equal(t, actual, web.Message{})
}
