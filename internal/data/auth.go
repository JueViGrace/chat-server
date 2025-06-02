package data

import (
	"context"
	"errors"
	"gps-tracker/internal/database"
	"gps-tracker/internal/types"

	"github.com/google/uuid"
)

type AuthStore interface {
	SignIn(r *types.SignInRequest) (res *types.AuthResponse, err error)
	SignUp(r *types.SignUpRequest) (res *types.AuthResponse, err error)
	Refresh(id uuid.UUID) (res *types.AuthResponse, err error)
	RecoverPassword(r *types.RecoverPasswordRequest) (msg string, err error)
	DeleteSession(id uuid.UUID) (err error)
	DeleteSessionByToken(token string) (err error)
}

func (s *storage) AuthStore() AuthStore {
	return NewAuthStore(s.ctx, s.queries)
}

type authStore struct {
	ctx     context.Context
	queries *database.Queries
}

func NewAuthStore(ctx context.Context, queries *database.Queries) AuthStore {
	return &authStore{
		ctx:     ctx,
		queries: queries,
	}
}

func (s *authStore) getSessionById(id uuid.UUID) (session *types.Session, err error) {
	session = new(types.Session)

	dbSession, err := s.queries.GetSessionById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	session, err = types.DbSessionToSession(&dbSession)
	if err != nil {
		return nil, err
	}

	return
}

func (s *authStore) SignIn(r *types.SignInRequest) (session *types.AuthResponse, err error) {
	user, err := s.queries.GetUser(s.ctx, database.GetUserParams{
		Email:    r.Email,
		Username: r.Email,
	})
	if err != nil {
		return nil, err
	}

	if user.DeletedAt.Valid {
		return nil, errors.New("this user was deleted")
	}

	userId, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	return
}

func (s *authStore) Refresh(id uuid.UUID) (session *types.AuthResponse, err error) {
	savedSession, err := s.getSessionById(id)
	if err != nil {
		return nil, err
	}

	session, err = createTokens(savedSession.ID)
	if err != nil {
		return nil, err
	}

	err = s.queries.UpdateSession(s.ctx, database.UpdateSessionParams{
		RefreshToken: session.RefreshToken,
		AccessToken:  session.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	return
}

func (s *authStore) DeleteSession(id uuid.UUID) (err error) {
	err = s.queries.DeleteSessionById(s.ctx, id.String())
	if err != nil {
		return err
	}

	return
}

func (s *authStore) DeleteSessionByToken(token string) (err error) {
	err = s.queries.DeleteSessionByToken(s.ctx, database.DeleteSessionByTokenParams{
		RefreshToken: token,
		AccessToken:  token,
	})
	if err != nil {
		return err
	}

	return
}

func createTokens(sessionId uuid.UUID) (*types.Session, error) {
	accessToken, err := types.CreateAccessToken(sessionId)
	if err != nil {
		return nil, err
	}

	refreshToken, err := types.CreateRefreshToken(sessionId)
	if err != nil {
		return nil, err
	}

	return &types.Session{
		ID:           sessionId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
