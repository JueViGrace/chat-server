package types

import (
	"errors"
	"gps-tracker/internal/database"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	Alias       string    `json:"alias"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	PhoneNumber string    `json:"phone_number"`
	BirthDate   string    `json:"birth_date"`
	Role        Role      `json:"-"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   string    `json:"-"`
}

type Role int

const (
	SystemUser Role = iota
	SystemAdmin
)

func (r Role) String() string {
	switch r {
	case SystemAdmin:
		return "Admin"
	case SystemUser:
		return "User"
	default:
		return "Unknown"
	}
}

func ParseRole(r int) (Role, error) {
	switch r {
	case int(SystemAdmin):
		return SystemAdmin, nil
	case int(SystemUser):
		return SystemUser, nil
	default:
		return 0, errors.New("invalid role")
	}
}

type UpdateUser struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Alias     string    `json:"alias" validate:"required"`
	BirthDate string    `json:"birth_date" validate:"required"`
}

type CheckUserFieldsRequest struct {
	UsedEmail       bool `json:"used_email"`
	UsedUsername    bool `json:"used_username"`
	UsedPhoneNumber bool `json:"used_phone_number"`
}

type ChangeEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type ChangeUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}

type ChangePhoneRequest struct {
	Username string `json:"username" validate:"required"`
}

func MapDbUser(db *database.User) (*User, error) {
	userId, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:          userId,
		FirstName:   db.Firstname,
		LastName:    db.Lastname,
		Username:    db.Username,
		Alias:       db.Alias.String,
		Email:       db.Email,
		PhoneNumber: db.PhoneNumber,
		BirthDate:   db.BirthDate,
		CreatedAt:   db.CreatedAt,
		UpdatedAt:   db.UpdatedAt,
	}, nil
}
