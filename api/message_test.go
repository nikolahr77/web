package api_test

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/web"
	"github.com/web/api"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) Get(id string) (web.Message, error) {
	args := m.Called(id)
	return args.Get(0).(web.Message), args.Error(1)
}

func (m *MockMessageRepository) Create(msg web.RequestMessage) (web.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(web.Message), args.Error(1)
}

func (m *MockMessageRepository) Update(id string, msg web.RequestMessage) (web.Message, error) {
	args := m.Called(id, msg)
	return args.Get(0).(web.Message), args.Error(1)
}

func (m *MockMessageRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateMessage(t *testing.T) {
	message := `{"name":"Hello", "content":"This is a hello message"}`
	req := httptest.NewRequest("POST", "/messages", strings.NewReader(message))
	w := httptest.NewRecorder()

	mr := web.RequestMessage{
		Name:    "Hello",
		Content: "This is a hello message",
	}

	m := web.Message{
		GUID:      "33123",
		Name:      "Hello",
		Content:   "This is a hello message",
		CreatedOn: time.Unix(15, 0).UTC(),
		UpdatedOn: time.Unix(25, 0).UTC(),
	}

	testObj := new(MockMessageRepository)

	testObj.On("Create", mr).Return(m, nil)

	r := mux.NewRouter()
	r.Handle("/messages", api.CreateMessage(testObj))
	r.ServeHTTP(w, req)
	actual := api.MessageDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.MessageDTO{
		GUID:      "33123",
		Name:      "Hello",
		Content:   "This is a hello message",
		CreatedOn: time.Unix(15, 0).UTC(),
		UpdatedOn: time.Unix(25, 0).UTC(),
	}

	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}

//func TestUpdateMessage(t *testing.T) {
//	message := `{"name":"Hello", "content":"This is a hello message"}`
//	req := httptest.NewRequest("POST", "/messages", strings.NewReader(message))
//	w := httptest.NewRecorder()
//
//
//	mr := web.RequestMessage{
//		Name: "Hello",
//		Content:"This is a hello message",
//	}
//
//	m := web.Message{
//		GUID: "33123",
//		Name: "Hello",
//		Content: "This is a hello message",
//		CreatedOn: time.Unix(15,0).UTC(),
//		UpdatedOn: time.Unix(25,0).UTC(),
//	}
//
//	testObj := new(MockMessageRepository)
//
//	testObj.On("Create", mr).Return(m,nil)
//
//	r := mux.NewRouter()
//	r.Handle("/messages", api.CreateMessage(testObj))
//	r.ServeHTTP(w,req)
//	actual := api.MessageDTO{}
//	json.NewDecoder(w.Body).Decode(&actual)
//	expected := api.MessageDTO{
//		GUID: "33123",
//		Name: "Hello",
//		Content: "This is a hello message",
//		CreatedOn: time.Unix(15,0).UTC(),
//		UpdatedOn: time.Unix(25,0).UTC(),
//	}
//
//	assert.Equal(t,expected,actual)
//
//	testObj.AssertExpectations(t)
//}
