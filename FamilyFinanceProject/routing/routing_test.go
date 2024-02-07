package routing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cursoGo/FamilyFinanceProject/common"
	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	r := gin.Default()
	CreateUsersRouting(r)

	// Define the test user data
	user := common.User{
		FirstName: "John",
		LastName:  "Doe",
		Relation:  "Family",
		Income:    1000.0,
		Expenses:  500.0,
		Balance:   500.0,
		Budget:    1000.0,
		Bill:      300.0,
	}

	// Convert the user data to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user to JSON: %v", err)
	}

	// Create a POST request with the user JSON data
	req, err := http.NewRequest("POST", "/users", strings.NewReader(string(userJSON)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	res := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(res, req)

	// Check the response status code
	if res.Code != http.StatusCreated {
		t.Errorf("Unexpected response status code: got %d, want %d", res.Code, http.StatusCreated)
	}

	// Parse the response body
	var createdUser common.User
	err = json.Unmarshal(res.Body.Bytes(), &createdUser)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Check the created user data
	if createdUser.FirstName != user.FirstName {
		t.Errorf("Unexpected created user first name: got %s, want %s", createdUser.FirstName, user.FirstName)
	}

	if createdUser.LastName != user.LastName {
		t.Errorf("Unexpected created user last name: got %s, want %s", createdUser.FirstName, user.FirstName)
	}
}

func TestGetAllUsers(t *testing.T) {
	r := gin.Default()
	CreateUsersRouting(r)

	// Populate users
	users = []common.User{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Relation:  "Family",
			Income:    1000.0,
			Expenses:  500.0,
			Balance:   500.0,
			Budget:    1000.0,
			Bill:      300.0,
		},
		// Add more users as needed
	}

	// Create a GET request
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to record the response
	res := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(res, req)

	// Check the response status code
	if res.Code != http.StatusOK {
		t.Errorf("Unexpected response status code: got %d, want %d", res.Code, http.StatusOK)
	}

	// Parse the response body
	var responseUsers []common.User
	err = json.Unmarshal(res.Body.Bytes(), &responseUsers)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Check the number of users returned
	if len(responseUsers) != len(users) {
		t.Errorf("Unexpected number of users returned: got %d, want %d", len(responseUsers), len(users))
	}

	// Check the user data
	for i, user := range users {
		responseUser := responseUsers[i]
		if responseUser.FirstName != user.FirstName {
			t.Errorf("Unexpected user first name: got %s, want %s", responseUser.FirstName, user.FirstName)
		}

	}
}
