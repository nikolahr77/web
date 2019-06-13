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
)

type MockContactRepository struct {
	mock.Mock
}

func (m *MockContactRepository) Get(id int64) (web.Contact, error) {
	args := m.Called(id)
	return args.Get(0).(web.Contact),args.Error(1)
}

func (m *MockContactRepository) Create(con web.Contact) (web.Contact, error){
	args := m.Called(con)
	return args.Get(0).(web.Contact),args.Error(1)
}

func (m *MockContactRepository) Delete(id int64) error{
	return nil
}

func (m *MockContactRepository) Update(id int64, con web.Contact) (web.Contact, error){
	return web.Contact{},nil
}

func TestCreateContact(t *testing.T) {
	contact := `{"name":"Ivan", "address":"Sofia", "age":23, "email":"test@test.com"}`
	req := httptest.NewRequest("POST", "/contacts", strings.NewReader(contact))
	w := httptest.NewRecorder()

	cr := web.RequestContact {
		Name: "Ivan",
		Address: "Sofia 1612",
		Age: 11,
		Email: "ivan@abv.bg",
	}

	c := web.Contact{
		Name: "Ivan",
		Address: "Sofia 1612",
		Age: 11,
		Email: "ivan@abv.bg",
	}

	testObj := new(MockContactRepository)

	testObj.On("Create",cr).Return(c,nil)

	r := mux.NewRouter()
	r.Handle("/contacts",api.CreateContact(testObj))
	r.ServeHTTP(w,req)
	actual := web.Contact{}
	expected := api.ContactDTO{
		ID:,

	}
	fmt.Printf("%#v\n", w.Body.String())
	json.NewDecoder(w.Body).Decode(&c)
	assert.Equal(t,cr,c)

	testObj.AssertExpectations(t)
}