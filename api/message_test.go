package api_test

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/web"
	"github.com/web/api"
	"net/http"
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
	assert.Equal(t, http.StatusOK, w.Code)

	testObj.AssertExpectations(t)
}

func TestCreateMessageReturnError(t *testing.T) {
	message := `{"name":"Hello", "content":"This is a hello message"}`
	req := httptest.NewRequest("POST", "/messages", strings.NewReader(message))
	w := httptest.NewRecorder()

	mr := web.RequestMessage{
		Name:    "Hello",
		Content: "This is a hello message",
	}

	testObj := new(MockMessageRepository)

	testObj.On("Create", mr).Return(web.Message{}, errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/messages", api.CreateMessage(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500

	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}

func TestCreateMessageMalformedJson(t *testing.T) {
	message := `{"name":334, "content":5532}`
	req := httptest.NewRequest("POST", "/messages", strings.NewReader(message))
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.Handle("/messages", api.CreateMessage(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 400
	assert.Equal(t, expected, actual)
}

func TestUpdateMessage(t *testing.T) {
	message := `{"name":"Hello", "content":"This is a hello message"}`
	req := httptest.NewRequest("POST", "/messages/1", strings.NewReader(message))
	w := httptest.NewRecorder()

	mr := web.RequestMessage{
		Name:    "Hello",
		Content: "This is a hello message",
	}

	m := web.Message{
		GUID:      "33123",
		Name:      "HellUpdate",
		Content:   "This is an updated message",
		CreatedOn: time.Unix(15, 0).UTC(),
		UpdatedOn: time.Unix(25, 0).UTC(),
	}

	testObj := new(MockMessageRepository)

	testObj.On("Update", "1", mr).Return(m, nil)

	r := mux.NewRouter()
	r.Handle("/messages/{id}", api.UpdateMessage(testObj))
	r.ServeHTTP(w, req)
	actual := api.MessageDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.MessageDTO{
		GUID:      "33123",
		Name:      "HellUpdate",
		Content:   "This is an updated message",
		CreatedOn: time.Unix(15, 0).UTC(),
		UpdatedOn: time.Unix(25, 0).UTC(),
	}

	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)

	testObj.AssertExpectations(t)
}

func TestUpdateMessageReturnError(t *testing.T) {
	message := `{"name":"Hello", "content":"This is a hello message"}`
	req := httptest.NewRequest("POST", "/messages/1", strings.NewReader(message))
	w := httptest.NewRecorder()

	mr := web.RequestMessage{
		Name:    "Hello",
		Content: "This is a hello message",
	}

	testObj := new(MockMessageRepository)

	testObj.On("Update", "1", mr).Return(web.Message{}, errors.New("Test Error"))

	r := mux.NewRouter()
	r.Handle("/messages/{id}", api.UpdateMessage(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500

	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}

func TestUpdateMessageMalformedJson(t *testing.T) {
	message := `{"name":41215}`
	req := httptest.NewRequest("PUT", "/messages/1", strings.NewReader(message))
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.Handle("/messages/1", api.UpdateMessage(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 400
	assert.Equal(t, expected, actual)
}

func TestDeleteMessage(t *testing.T) {
	message := `{"name":"Hello", "content":"This is a hello message"}`
	req := httptest.NewRequest("POST", "/messages/1", strings.NewReader(message))
	w := httptest.NewRecorder()

	testObj := new(MockMessageRepository)

	testObj.On("Delete", "1").Return(nil)

	r := mux.NewRouter()
	r.Handle("/messages/{id}", api.DeleteMessage(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 200

	assert.Equal(t, expected, actual)
}

func TestDeleteMessageReturnError(t *testing.T) {
	message := `{"name":"Hello", "content":"This is a hello message"}`
	req := httptest.NewRequest("POST", "/messages/1", strings.NewReader(message))
	w := httptest.NewRecorder()

	testObj := new(MockMessageRepository)

	testObj.On("Delete", "1").Return(errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/messages/{id}", api.DeleteMessage(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500

	assert.Equal(t, expected, actual)
}

func TestGetMessage(t *testing.T) {
	message := `{"name":"Test","content":"Getting test message"}`
	req := httptest.NewRequest("GET", "/messages/1", strings.NewReader(message))
	w := httptest.NewRecorder()

	m := web.Message{
		GUID:      "ab33",
		Name:      "Test",
		Content:   "Getting test message",
		CreatedOn: time.Unix(15, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}

	testObj := new(MockMessageRepository)

	testObj.On("Get", "1").Return(m, nil)

	r := mux.NewRouter()
	r.Handle("/messages/{id}", api.GetMessage(testObj))
	r.ServeHTTP(w, req)
	actual := api.MessageDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.MessageDTO{
		GUID:      "ab33",
		Name:      "Test",
		Content:   "Getting test message",
		CreatedOn: time.Unix(15, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)

	testObj.AssertExpectations(t)
}

func TestGetMessageReturnError(t *testing.T) {
	message := `{"name":"Test","Content":"Getting test message"}`
	req := httptest.NewRequest("GET", "/messages/1", strings.NewReader(message))
	w := httptest.NewRecorder()

	testObj := new(MockMessageRepository)

	testObj.On("Get", "1").Return(web.Message{}, errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/messages/{id}", api.GetMessage(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}
