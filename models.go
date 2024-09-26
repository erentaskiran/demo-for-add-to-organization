package main

import "time"

type Organization struct {
	UserID      string    `json:"user_id"`
	Name        string    `json:"organization_name"`
	Description string    `json:"organization_description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type OrganizationBody struct {
	OrganizationName        string `json:"organization_name"`
	OrganizationDescription string `json:"organization_description"`
}

type OrganizationInfo struct {
	OrganizationId          string
	OrganizationName        string
	OrganizationDescription string
}

type OrganizationCreated struct {
	Id          string
	Name        string
	Description string
	CreatedBy   string
}
type OrganizationUserCreated struct {
	OrganizationId string
	UserId         string
	Role           string
}

type AddUserOrganizationPayload struct {
	OrganizationID string `json:"organization_id"`
	UserID         string `json:"user_id"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}
type User struct {
	Id        string    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
