package model

type User struct {
	Name     string `json:"name"`
	Surname  string `validate:"required,surname,omitempty"`
	Password string `json:"password"`
}
