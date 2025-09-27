package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

// Mock Repository function will be simualates the data
// In case in order the test the user services
func (m *MockRepo) GetUser(id int) (*User, error) {
	args := m.Called(id)
	if user := args.Get(0); user != nil {
		return user.(*User), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetUserName_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	service := UserService{repo: mockRepo}

	mockRepo.On("GetUser", 1).Return(&User{ID: 1, Name: "John Doe"}, nil)

	// Call the functions in the Services
	name, err := service.GetUserName(1)

	// Asserations
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", name)

	// Verify that mocks was used as expected
	mockRepo.AssertExpectations(t)
}

func TestGetUserName_Error(t *testing.T) {
	mockRepo := new(MockRepo)
	service := UserService{repo: mockRepo}
	mockRepo.On("GetUser", 2).Return(nil, errors.New("user not found"))

	name, err := service.GetUserName(2)

	assert.Error(t, err)
	assert.Equal(t, "test", name)
	mockRepo.AssertExpectations(t)

}

/*
	1. Slice has the dynamically resizeable based on the elements
	which we are appending in the code as well and usually make
	a better copies to allocates for the memories.
	2.

*/
