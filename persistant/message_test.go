package persistant_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/persistant"
	"testing"
)

func TestCreateMessageRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}

	mr := persistant.NewMessageRepository(DB, clock)

	newMsg := web.RequestMessage{
		Name:    "TestMSG",
		Content: "This is a test message",
	}
	userID := uuid.New()
	actual, err := mr.Create(newMsg, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Message{
		GUID:      actual.GUID,
		Name:      "TestMSG",
		Content:   "This is a test message",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
		UserID:    userID.String(),
	}

	assert.Equal(t, expected, actual)
	dbCleaner(DB, "messages")
}

func TestUpdateMessageRepository(t *testing.T) {
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
	actual, err := mr.Update(old.GUID, newMsg, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Message{
		GUID:      actual.GUID,
		Name:      "NewMSG",
		Content:   "This is the new test message",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
	}

	assert.Equal(t, expected, actual)
	dbCleaner(DB, "messages")
}

func TestDeleteMessageRepository(t *testing.T) {
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

	assert.Equal(t, err, nil)
}

func TestGetMessageRepository(t *testing.T) {
	clock := fakeClock{
		Seconds: 25000,
	}

	mr := persistant.NewMessageRepository(DB, clock)

	newMsg := web.RequestMessage{
		Name:    "NewMSG",
		Content: "This is the new test message",
	}
	userID := uuid.New()
	contact, err := mr.Create(newMsg, userID.String())
	if err != nil {
		panic(err)
	}
	actual, err := mr.Get(contact.GUID, userID.String())
	if err != nil {
		panic(err)
	}

	expected := web.Message{
		GUID:      actual.GUID,
		Name:      "NewMSG",
		Content:   "This is the new test message",
		CreatedOn: actual.CreatedOn, //I should't do this
		UpdatedOn: actual.UpdatedOn,
		UserID:    contact.UserID,
	}

	assert.Equal(t, expected, actual)
	dbCleaner(DB, "messages")
}
