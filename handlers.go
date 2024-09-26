package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Handler struct {
	Db  *sql.DB
	Cfg *aws.Config
}

func newHandler(db *sql.DB, cfg *aws.Config) *Handler {
	return &Handler{
		Db:  db,
		Cfg: cfg,
	}
}
func (h *Handler) HandleUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		h.HandleCreateUser(w, r)
	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
}
func (h *Handler) HandleOrganization(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.HandleCreateOrganization(w, r)
	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
}
func (h *Handler) HandleOrganizationUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.HandleAddOrganizationUser(w, r)
	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user := NewUserRepository(h.Db)

	result, err := user.CreateUser(&payload)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(w, http.StatusOK, result)
}

func (h *Handler) HandleCreateOrganization(w http.ResponseWriter, r *http.Request) {
	var payload Organization
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	organization := NewOrganizationRepository(h.Db)
	org := &Organization{
		Name:        payload.Name,
		Description: payload.Description,
		CreatedBy:   payload.UserID,
	}
	result, err := organization.CreateOrganization(org)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	JSONResponse(w, http.StatusOK, result)
}

func (h *Handler) HandleAddOrganizationUser(w http.ResponseWriter, r *http.Request) {
	var payload AddUserOrganizationPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	u := NewUserRepository(h.Db)
	organization := NewOrganizationRepository(h.Db)
	role, err := organization.GetOrganizationUser(payload.UserID, payload.OrganizationID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		JSONError(w, http.StatusInternalServerError, "Unauthoruized")
		return
	}
	if role == "admin" || role == "owner" {
		user, err := u.GetUserWithEmail(payload.Email)
		if err != nil {
			helper := NewHelper(h.Cfg)
			err := helper.SendEmail(payload.Email)
			if err != nil {
				JSONError(w, http.StatusInternalServerError, err.Error())
				return
			}
			JSONResponse(w, http.StatusOK, "email sent succesfully")
		} else {

			organization_user := &OrganizationUserCreated{
				OrganizationId: payload.OrganizationID,
				UserId:         user.Id,
				Role:           payload.Role,
			}
			result, err := organization.CreateOrganizationUser(organization_user)
			if err != nil {
				JSONError(w, http.StatusInternalServerError, err.Error())
				return
			}
			JSONResponse(w, http.StatusOK, result)
		}
	}
}
