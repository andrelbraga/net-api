package entities

import (
	"errors"
	"net/http"
)

type User struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserRequest struct {
	*User
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.User == nil {
		return errors.New("missing required User fields.")
	}
	return nil
}
