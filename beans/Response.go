package beans

import (
	"echoPoc/types"
)

type List1Response struct {
	Httpstatus bool `json:"success"`
	Data  *types.User `json:"data"`
}
type ListallResponse struct {
	Httpstatus bool `json:"success"`
	Data  *types.Users `json:"data"`
}
//func (List1Response) Render (w http.ResponseWriter, r *http.Request) error{
//	return nil
//}

func List1user(user *types.User) *List1Response {
	resp := &List1Response{Data: user}
	return resp
}
func Listalluser(users *types.Users) *ListallResponse {
	resp := &ListallResponse{Data: users}
	return resp
}