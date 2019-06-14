package api_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/web"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) Get(id string) (web.Message,error){
	args := m.Called(id)
	return args.Get(0).(web.Message), args.Error(1)
}

func (m *MockMessageRepository) Create(msg web.RequestMessage) (web.Message,error){
	args := m.Called(msg)
	return args.Get(0).(web.Message), args.Error(1)
}

func (m *MockMessageRepository) Update(id string,msg web.RequestMessage) (web.Message,error){
	args := m.Called(id,msg)
	return args.Get(0).(web.Message),args.Error(1)
}

func (m *MockMessageRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
