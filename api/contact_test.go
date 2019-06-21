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

type MockContactRepository struct {
	mock.Mock
}

func (m *MockContactRepository) Get(id string) (web.Contact, error) {
	args := m.Called(id)
	return args.Get(0).(web.Contact), args.Error(1)
}

func (m *MockContactRepository) Create(con web.RequestContact) (web.Contact, error) {
	args := m.Called(con)
	return args.Get(0).(web.Contact), args.Error(1)
}

func (m *MockContactRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockContactRepository) Update(id string, con web.RequestContact) (web.Contact, error) {
	args := m.Called(id, con)
	return args.Get(0).(web.Contact), args.Error(1)
}

func TestCreateContact(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("POST", "/contacts", strings.NewReader(contact))
	w := httptest.NewRecorder()

	cr := web.RequestContact{
		Name:    "Ivan",
		Address: "Sofia 1612",
		Age:     23,
		Email:   "test@test.com",
	}

	c := web.Contact{
		GUID:      "512341",
		Name:      "Ivan",
		Address:   "Sofia 1612",
		Age:       23,
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}

	testObj := new(MockContactRepository)

	testObj.On("Create", cr).Return(c, nil)

	r := mux.NewRouter()
	r.Handle("/contacts", api.CreateContact(testObj))
	r.ServeHTTP(w, req)
	actual := api.ContactDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.ContactDTO{
		GUID:      "512341",
		Name:      "Ivan",
		Age:       23,
		Address:   "Sofia 1612",
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)
	testObj.AssertExpectations(t)
}

func TestCreateContactError(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("POST", "/contacts", strings.NewReader(contact))
	w := httptest.NewRecorder()

	cr := web.RequestContact{
		Name:    "Ivan",
		Address: "Sofia 1612",
		Age:     23,
		Email:   "test@test.com",
	}

	testObj := new(MockContactRepository)

	testObj.On("Create", cr).Return(web.Contact{}, errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/contacts", api.CreateContact(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}

func TestCreateContactMalformedJson(t *testing.T) {
	contact := `{fffsgffa}`
	req := httptest.NewRequest("POST", "/contacts", strings.NewReader(contact))

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Handle("/contacts", api.CreateContact(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 400
	assert.Equal(t, expected, actual)
}

func TestGetContact(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("GET", "/contacts/1", strings.NewReader(contact))
	w := httptest.NewRecorder()

	c := web.Contact{
		GUID:      "1",
		Name:      "Ivan",
		Address:   "Sofia 1612",
		Age:       23,
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}

	testObj := new(MockContactRepository)

	testObj.On("Get", "1").Return(c, nil)

	r := mux.NewRouter()
	r.Handle("/contacts/{id}", api.GetContact(testObj))
	r.ServeHTTP(w, req)
	actual := api.ContactDTO{}

	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.ContactDTO{
		GUID:      "1",
		Name:      "Ivan",
		Age:       23,
		Address:   "Sofia 1612",
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)
	testObj.AssertExpectations(t)
}

func TestGetContactReturnError(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("GET", "/contacts/1", strings.NewReader(contact))
	w := httptest.NewRecorder()

	testObj := new(MockContactRepository)

	testObj.On("Get", "1").Return(web.Contact{}, errors.New("Test Error"))

	r := mux.NewRouter()
	r.Handle("/contacts/{id}", api.GetContact(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)
	testObj.AssertExpectations(t)
}

func TestUpdateContact(t *testing.T) {
	contact := `{"name":"Stefan", "address":"Plovdiv", "age":34, "email":"stef@test.com"}`
	req := httptest.NewRequest("POST", "/contacts/1", strings.NewReader(contact))
	w := httptest.NewRecorder()

	cr := web.RequestContact{
		Name:    "Stefan",
		Address: "Plovdiv",
		Age:     34,
		Email:   "stef@test.com",
	}

	c := web.Contact{
		GUID:      "512341",
		Name:      "Ivan",
		Address:   "Sofia 1612",
		Age:       23,
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}

	testObj := new(MockContactRepository)

	testObj.On("Update", "1", cr).Return(c, nil)

	r := mux.NewRouter()
	r.Handle("/contacts/{id}", api.UpdateContact(testObj))
	r.ServeHTTP(w, req)
	actual := api.ContactDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.ContactDTO{
		GUID:      "512341",
		Name:      "Ivan",
		Age:       23,
		Address:   "Sofia 1612",
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0).UTC(),
		UpdatedOn: time.Unix(20, 0).UTC(),
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)
	testObj.AssertExpectations(t)
}

func TestUpdateContactReturnError(t *testing.T) {
	contact := `{"name":"Stefan", "address":"Plovdiv", "age":34, "email":"stef@test.com"}`
	req := httptest.NewRequest("POST", "/contacts/1", strings.NewReader(contact))
	w := httptest.NewRecorder()

	cr := web.RequestContact{
		Name:    "Stefan",
		Address: "Plovdiv",
		Age:     34,
		Email:   "stef@test.com",
	}

	testObj := new(MockContactRepository)

	testObj.On("Update", "1", cr).Return(web.Contact{}, errors.New("Test Error"))

	r := mux.NewRouter()
	r.Handle("/contacts/{id}", api.UpdateContact(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)
	testObj.AssertExpectations(t)
}

func TestUpdateContactMalformedJson(t *testing.T) {
	contact := `{"name":33931}`
	req := httptest.NewRequest("PUT", "/contacts/1", strings.NewReader(contact))
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Handle("/contacts/1", api.UpdateContact(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 400
	assert.Equal(t, expected, actual)
}

func TestDeleteContact(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("DELETE", "/contacts/512341", strings.NewReader(contact))
	w := httptest.NewRecorder()

	testObj := new(MockContactRepository)

	testObj.On("Delete", "512341").Return(nil)

	r := mux.NewRouter()
	r.Handle("/contacts/{id}", api.DeleteContact(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 200
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}

func TestDeleteContactReturnError(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("DELETE", "/contacts/512341", strings.NewReader(contact))
	w := httptest.NewRecorder()

	testObj := new(MockContactRepository)

	testObj.On("Delete", "512341").Return(errors.New("Test Error"))

	r := mux.NewRouter()
	r.Handle("/contacts/{id}", api.DeleteContact(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}
