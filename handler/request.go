package handler

import (
	"fmt"
)

func errParamIsRequired(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateOpeningRequest struct {
	Role        string `json:"role"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Salary      int64  `json:"salary"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Description == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is malformed or empty")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role        string `json:"role"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Salary      int64  `json:"salary"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validate it
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Description != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}

	// If no field is provided, return an error
	return fmt.Errorf("at least one field is required to update the opening")
}
