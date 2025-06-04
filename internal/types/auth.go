package types

import (
	"database/sql"
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
	Role         Role
}

type AuthResponse struct {
	ID           uuid.UUID `json:"id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Alias       string `json:"alias"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type RecoverPasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

func DbSessionToSession(db *database.Session) (*Session, error) {
	id, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(db.UserID)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:           id,
		AccessToken:  db.AccessToken,
		RefreshToken: db.RefreshToken,
		UserID:       userId,
	}, nil
}

func CreateUser(r *SignUpRequest) (*database.CreateUserParams, error) {
	userId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	encPass, err := HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	return &database.CreateUserParams{
		ID:        userId.String(),
		Firstname: r.FirstName,
		Lastname:  r.LastName,
		Username:  r.Username,
		Alias: sql.NullString{
			String: r.Alias,
			Valid:  true,
		},
		Email:       r.Email,
		Password:    encPass,
		PhoneNumber: r.PhoneNumber,
		Role:        int64(SystemUser),
	}, nil
}
