package types

import (
	"github.com/Kamva/mgm"
 "github.com/go-playground/validator/v10"

)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name      string `json:"name" validate:"required" bson:"name"`
	Email	string	`json:"email" validate:"required,email" bson:"email"`
	Age 	int 	`json:"age" validate:"required" bson:"age"`
}
type Users []User

type UserValidator struct {
	Validator *validator.Validate
}
func (u *UserValidator) Validate(i interface{}) error {
	return u.Validator.Struct(i)
}
