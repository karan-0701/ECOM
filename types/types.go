package types

import "time"

type UserStore interface {
	GetUseByEmail(email string) (*User, error)
	GetUseByID(id int) (*User, error)
	CreateUser(User) error
}
type RegisteredUserPayload struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Email     string `json: "email"`
	Password  string `json:"firstName"`
}

type User struct {
	ID        int       `json: "id"`
	FirstName string    `json: "firstName"`
	LastName  string    `json: "lastName"`
	Email     string    `json: "email"`
	Password  string    `json: "-"`
	CreatedAt time.Time `json: "createdAt"`
}
