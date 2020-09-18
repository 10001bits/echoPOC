package types

import (
	"github.com/Kamva/mgm"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name      string `json:"name" bson:"name"`
	Email	string	`json:"email" bson:"email"`
	Age 	int 	`json:"age" bson:"age"`
}
type Users []User