package api_test

import (
	"encoding/json"
	"errors"
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
		CreatedOn: time.Unix(10, 0),
		UpdatedOn: time.Unix(20, 0),
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
		CreatedOn: time.Unix(10, 0),
		UpdatedOn: time.Unix(20, 0),
	}
	assert.Equal(t, expected, actual)

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

//func TestCreateContactErrorJson(t *testing.T) {
//	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
//	req := httptest.NewRequest("POST", "/contacts", strings.NewReader(contact))
//	w := httptest.NewRecorder()
//
//	cr := web.RequestContact{
//		Name:    "Ivan",
//		Address: "Sofia 1612",
//		Age:     23,
//		Email:   "test@test.com",
//	}
//
//	c := web.Contact{
//		GUID:      "512341",
//		Name:      "Ivan",
//		Address:   "Sofia 1612",
//		Age:       23,
//		Email:     "test@test.com",
//		CreatedOn: time.Unix(10, 0),
//		UpdatedOn: time.Unix(20, 0),
//	}
//
//	testObj := new(MockContactRepository)
//
//	testObj.On("json.NewDecoder", ).Return(c, errors.New("test error"))
//
//	r := mux.NewRouter()
//	r.Handle("/contacts", api.CreateContact(testObj))
//	r.ServeHTTP(w, req)
//	actual := w.Code
//	expected := 500
//	assert.Equal(t, expected, actual)
//
//	testObj.AssertExpectations(t)
//}

func TestGetContact(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia 1612", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("POST", "/contacts/1", strings.NewReader(contact))
	w := httptest.NewRecorder()

	id := 1

	c := web.Contact{
		GUID:      "1",
		Name:      "Ivan",
		Address:   "Sofia 1612",
		Age:       23,
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0),
		UpdatedOn: time.Unix(20, 0),
	}

	testObj := new(MockContactRepository)

	testObj.On("Get", id).Return(c, nil)

	r := mux.NewRouter()
	r.Handle("/contacts/", api.GetContact(testObj))
	r.ServeHTTP(w, req)
	actual := api.ContactDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.ContactDTO{
		GUID:      "1",
		Name:      "Ivan",
		Age:       23,
		Address:   "Sofia 1612",
		Email:     "test@test.com",
		CreatedOn: time.Unix(10, 0),
		UpdatedOn: time.Unix(20, 0),
	}
	assert.Equal(t, expected, actual)

	testObj.AssertExpectations(t)
}
