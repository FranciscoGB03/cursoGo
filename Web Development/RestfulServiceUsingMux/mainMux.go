package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var users []User

func main() {
	router := mux.NewRouter()

	//Add sample data
	users = append(users, User{ID: "1", Name: "Sivaprasad", Email: "sivaprasad@example.com", Password: "password"})
	users = append(users, User{ID: "2", Name: "Tamatam", Email: "tamatam@example.com", Password: "password"})

	// Define routes
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8001", router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, user := range users {
		if user.ID == params["id"] {
			users[i] = User{
				ID:       params["id"],
				Name:     user.Name,
				Email:    user.Email,
				Password: user.Password,
			}
			_ = json.NewDecoder(r.Body).Decode(&users[i])
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, user := range users {
		if user.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
