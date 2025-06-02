package types

import "github.com/google/uuid"

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
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   string    `json:"-"`
}

type UpdateUser struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Alias     string    `json:"alias" validate:"required"`
	BirthDate string    `json:"birth_date" validate:"required"`
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
