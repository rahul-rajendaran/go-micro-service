package service

import "errors"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: "1", Name: "Alice", Email: "alice@example.com"},
	{ID: "2", Name: "Bob", Email: "bob@example.com"},
}

func GetAllUsers() []User {
	return users
}

func CreateUser(u User) User {
	u.ID = string(len(users) + 1)
	users = append(users, u)
	return u
}

func GetUserByID(id string) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, errors.New("not found")
}
