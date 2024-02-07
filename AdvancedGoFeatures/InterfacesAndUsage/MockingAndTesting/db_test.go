package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Query(query string) ([]string, error) {
	args := m.Called(query)
	return args.Get(0).([]string), args.Error(1)
}

func TestGetDataFromDatabase(t *testing.T) {
	mockDB := new(MockDatabase)
	mockResults := []string{"user1", "user2", "user3"}

	//Expect Query method to be called with "SELECT* FROM users"
	mockDB.On("Query", "SELECT * FROM users").Return(mockResults, nil)

	result, err := GetDataFromDatabase(mockDB)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResults := mockResults

	if !reflect.DeepEqual(result, expectedResults) {
		t.Errorf("expected %v, got %v", expectedResults, result)
	}

	mockDB.AssertExpectations(t)
}
