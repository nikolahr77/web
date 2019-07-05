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

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Get(guid string) (web.User, error) {
	args := m.Called(guid)
	return args.Get(0).(web.User), args.Error(1)
}

func (m *MockUserRepository) GetByName(name string) (web.User, error) {
	args := m.Called(name)
	return args.Get(0).(web.User), args.Error(1)
}

func (m *MockUserRepository) Create(usr web.RequestUser) (web.User, error) {
	args := m.Called(usr)
	return args.Get(0).(web.User), args.Error(1)
}

func (m *MockUserRepository) Update(guid string, usr web.RequestUser) (web.User, error) {
	args := m.Called(guid, usr)
	return args.Get(0).(web.User), args.Error(1)
}

func (m *MockUserRepository) Delete(guid string) error {
	args := m.Called(guid)
	return args.Error(0)

}

func TestCreateUser(t *testing.T) {
	user := `{"name":"Toncho Tonchev","age":43,"password":"sss1234","email":"ton@gmail.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(user))
	w := httptest.NewRecorder()

	ur := web.RequestUser{
		Name:     "Toncho Tonchev",
		Age:      43,
		Password: "sss1234",
		Email:    "ton@gmail.com",
	}

	usr := web.User{
		GUID:      "71-4b3-a32",
		Name:      "Toncho Tonchev",
		Age:       43,
		Password:  "sss1234",
		Email:     "ton@gmail.com",
		CreatedOn: time.Unix(80, 0).UTC(),
		UpdatedOn: time.Unix(80, 0).UTC(),
	}
	testObj := new(MockUserRepository)

	testObj.On("Create", ur).Return(usr, nil)

	r := mux.NewRouter()
	r.Handle("/users", api.CreateUser(testObj))
	r.ServeHTTP(w, req)
	actual := api.UserDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.UserDTO{
		GUID:      "71-4b3-a32",
		Name:      "Toncho Tonchev",
		Age:       43,
		Password:  "sss1234",
		Email:     "ton@gmail.com",
		CreatedOn: time.Unix(80, 0).UTC(),
		UpdatedOn: time.Unix(80, 0).UTC(),
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)
	testObj.AssertExpectations(t)

}

func TestCreateUserReturnError(t *testing.T) {
	user := `{"name":"Toncho Tonchev","age":43,"password":"sss1234","email":"ton@gmail.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(user))
	w := httptest.NewRecorder()

	ur := web.RequestUser{
		Name:     "Toncho Tonchev",
		Age:      43,
		Password: "sss1234",
		Email:    "ton@gmail.com",
	}

	testObj := new(MockUserRepository)

	testObj.On("Create", ur).Return(web.User{}, errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/users", api.CreateUser(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)
	testObj.AssertExpectations(t)

}

func TestCreateUserMalformedJson(t *testing.T) {
	user := `{"name":521521,"age":43,"p215234","email55":"ton@g5215mail.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(user))
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.Handle("/users", api.CreateUser(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	json.NewDecoder(w.Body).Decode(&actual)
	expected := 400
	assert.Equal(t, expected, actual)

}

func TestUpdateUser(t *testing.T) {
	user := `{"name":"Ivcho","age":13,"password":"parola313","email":"iv@gmail.com"}`
	req := httptest.NewRequest("POST", "/users/1", strings.NewReader(user))
	w := httptest.NewRecorder()

	ur := web.RequestUser{
		Name:     "Ivcho",
		Age:      13,
		Password: "parola313",
		Email:    "iv@gmail.com",
	}

	usr := web.User{
		GUID:      "71-4b3-a32",
		Name:      "Toncho Tonchev",
		Age:       43,
		Password:  "sss1234",
		Email:     "ton@gmail.com",
		CreatedOn: time.Unix(80, 0).UTC(),
		UpdatedOn: time.Unix(120, 0).UTC(),
	}
	testObj := new(MockUserRepository)

	testObj.On("Update", "1", ur).Return(usr, nil)

	r := mux.NewRouter()
	r.Handle("/users/{id}", api.UpdateUser(testObj))
	r.ServeHTTP(w, req)
	actual := api.UserDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.UserDTO{
		GUID:      "71-4b3-a32",
		Name:      "Toncho Tonchev",
		Age:       43,
		Password:  "sss1234",
		Email:     "ton@gmail.com",
		CreatedOn: time.Unix(80, 0).UTC(),
		UpdatedOn: time.Unix(120, 0).UTC(),
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)
	testObj.AssertExpectations(t)

}

func TestUpdateUserReturnError(t *testing.T) {
	user := `{"name":"Ivcho","age":13,"password":"parola313","email":"iv@gmail.com"}`
	req := httptest.NewRequest("POST", "/users/1", strings.NewReader(user))
	w := httptest.NewRecorder()

	ur := web.RequestUser{
		Name:     "Ivcho",
		Age:      13,
		Password: "parola313",
		Email:    "iv@gmail.com",
	}

	testObj := new(MockUserRepository)

	testObj.On("Update", "1", ur).Return(web.User{}, errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/users/{id}", api.UpdateUser(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)
	testObj.AssertExpectations(t)

}

func TestUpdateUserMalformedJson(t *testing.T) {
	user := `{"name"251"Ivcho","age125":13,"passwor"parola313","emailgmail.com"}`
	req := httptest.NewRequest("POST", "/users/1", strings.NewReader(user))
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.Handle("/users/{id}", api.UpdateUser(nil))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 400
	assert.Equal(t, expected, actual)

}

func TestDeleteUser(t *testing.T) {
	user := `{"name":"Ivcho","age":13,"password":"parola313","email":"iv@gmail.com"}`
	req := httptest.NewRequest("POST", "/users/1", strings.NewReader(user))
	w := httptest.NewRecorder()

	testObj := new(MockUserRepository)

	testObj.On("Delete", "1").Return(nil)

	r := mux.NewRouter()
	r.Handle("/users/{id}", api.DeleteUser(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 200
	assert.Equal(t, expected, actual)
	testObj.AssertExpectations(t)

}

func TestDeleteUserReturnError(t *testing.T) {
	user := `{"name":"Ivcho","age":13,"password":"parola313","email":"iv@gmail.com"}`
	req := httptest.NewRequest("POST", "/users/1", strings.NewReader(user))
	w := httptest.NewRecorder()

	testObj := new(MockUserRepository)

	testObj.On("Delete", "1").Return(errors.New("test eror"))

	r := mux.NewRouter()
	r.Handle("/users/{id}", api.DeleteUser(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)
	testObj.AssertExpectations(t)

}

func TestGetUser(t *testing.T) {
	user := `{"name":"Ivcho","age":13,"password":"parola313","email":"iv@gmail.com"}`
	req := httptest.NewRequest("POST", "/users/41-b3", strings.NewReader(user))
	w := httptest.NewRecorder()

	usr := web.User{
		GUID:      "41-b3",
		Name:      "Toncho Tonchev",
		Age:       43,
		Password:  "sss1234",
		Email:     "ton@gmail.com",
		CreatedOn: time.Unix(80, 0).UTC(),
		UpdatedOn: time.Unix(120, 0).UTC(),
	}

	testObj := new(MockUserRepository)

	testObj.On("Get", "41-b3").Return(usr, nil)

	r := mux.NewRouter()
	r.Handle("/users/{id}", api.GetUser(testObj))
	r.ServeHTTP(w, req)
	actual := api.UserDTO{}
	json.NewDecoder(w.Body).Decode(&actual)
	expected := api.UserDTO{
		GUID:      "41-b3",
		Name:      "Toncho Tonchev",
		Age:       43,
		Password:  "sss1234",
		Email:     "ton@gmail.com",
		CreatedOn: time.Unix(80, 0).UTC(),
		UpdatedOn: time.Unix(120, 0).UTC(),
	}
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, w.Code)
	testObj.AssertExpectations(t)
}

func TestGetUserReturnError(t *testing.T) {
	user := `{"name":"Ivcho","age":13,"password":"parola313","email":"iv@gmail.com"}`
	req := httptest.NewRequest("POST", "/users/41-b3", strings.NewReader(user))
	w := httptest.NewRecorder()

	testObj := new(MockUserRepository)

	testObj.On("Get", "41-b3").Return(web.User{}, errors.New("test error"))

	r := mux.NewRouter()
	r.Handle("/users/{id}", api.GetUser(testObj))
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := 500
	assert.Equal(t, expected, actual)
	testObj.AssertExpectations(t)
}
