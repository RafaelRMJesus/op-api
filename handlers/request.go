package handler

import (
	"fmt"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsReq(name, typ string) error {
	return fmt.Errorf("param %s is required and must be %s", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsReq("role", "string")
	}
	if r.Company == "" {
		return errParamIsReq("company", "string")
	}
	if r.Location == "" {
		return errParamIsReq("location", "string")
	}
	if r.Remote == nil {
		return errParamIsReq("remote", "bool")
	}
	if r.Link == "" {
		return errParamIsReq("link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsReq("salary", "int64")
	}
	return nil
}
