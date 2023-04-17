package model

type User struct {
	Name     string `json:"name"`
	Surname  string `validate:"required,surname,omitempty"`
	Password string `validate:"required,gte=7,lte=130,omitempty"`
}
