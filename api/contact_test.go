package api_test

import (
	"encoding/json"
	"fmt"
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
	return nil
}

func (m *MockContactRepository) Update(id string, con web.RequestContact) (web.Contact, error) {
	return web.Contact{}, nil
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
	fmt.Printf("%#v\n", w.Body.String())
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
	fmt.Printf("%#v\n", w.Body.String())
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
