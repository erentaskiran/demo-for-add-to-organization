package main

import (
	"database/sql"
	"time"
)

type OrganizationUserRepository struct {
	db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) *OrganizationUserRepository {
	return &OrganizationUserRepository{db: db}
}
func (r *OrganizationUserRepository) CreateOrganizationUser(user *OrganizationUserCreated) (OrganizationUserCreated, error) {
	query := `
		INSERT INTO organization_users (organization_id, user_id, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING organization_id, user_id, role
	`
	now := time.Now()

	var createdOrganizationUser OrganizationUserCreated
	err := r.db.QueryRow(
		query,
		user.OrganizationId,
		user.UserId,
		user.Role,
		now,
		now,
	).Scan(
		&createdOrganizationUser.OrganizationId,
		&createdOrganizationUser.UserId,
		&createdOrganizationUser.Role,
	)

	if err != nil {
		return createdOrganizationUser, err
	}

	return createdOrganizationUser, nil
}

func (r *OrganizationUserRepository) GetOrganizationUser(user_id, organization_id string) (string, error) {
	query := `
		SELECT role FROM organization_users where organization_id = $1 and user_id = $2
	`

	var role string
	err := r.db.QueryRow(query, organization_id, user_id).Scan(&role)

	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *OrganizationUserRepository) CreateOrganization(organization *Organization) (OrganizationCreated, error) {
	query := `
		INSERT INTO organizations (name, description, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, description, name, created_by
	`
	now := time.Now()

	var createdOrganization OrganizationCreated
	err := r.db.QueryRow(
		query,
		organization.Name,
		organization.Description,
		organization.CreatedBy,
		now,
		now,
	).Scan(
		&createdOrganization.Id,
		&createdOrganization.Name,
		&createdOrganization.Description,
		&createdOrganization.CreatedBy,
	)

	if err != nil {
		return createdOrganization, err
	}

	return createdOrganization, nil
}
