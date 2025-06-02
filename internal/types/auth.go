package types

import (
	"gps-tracker/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthDataHandler = func(*fiber.Ctx, *Session) error

type Session struct {
	ID           uuid.UUID
	AccessToken  string
	RefreshToken string
	UserID       uuid.UUID
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Username    string `json:"username"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	BirthDate   string `json:"birth_date" validate:"required"`
}

type RecoverPasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

func DbSessionToSession(db *database.Session) (*Session, error) {
	id, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:           id,
		AccessToken:  db.AccessToken,
		RefreshToken: db.RefreshToken,
	}, nil
}
