package data

import (
	"context"
	"errors"
	"gps-tracker/internal/database"
	"gps-tracker/internal/types"

	"github.com/google/uuid"
)

type AuthStore interface {
	GetSessionById(id uuid.UUID) (session *types.Session, err error)
	LogIn(r *types.SignInRequest) (tokens *types.AuthResponse, err error)
	SignUp(r *types.SignUpRequest) (tokens *types.AuthResponse, err error)
	Refresh(session *types.Session) (tokens *types.AuthResponse, err error)
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

func (s *authStore) GetSessionById(id uuid.UUID) (session *types.Session, err error) {
	session = new(types.Session)

	dbSession, err := s.queries.GetSessionById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	user, err := s.queries.GetUserById(s.ctx, dbSession.UserID)
	if err != nil {
		return nil, err
	}

	userRole, err := types.ParseRole(int(user.Role))
	if err != nil {
		return nil, err
	}

	session, err = types.DbSessionToSession(&dbSession)
	if err != nil {
		return nil, err
	}

	session.Role = userRole

	return
}

func (s *authStore) LogIn(r *types.SignInRequest) (*types.AuthResponse, error) {
	tokens := new(types.AuthResponse)
	session := new(types.Session)

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

	if !types.ValidatePassword(r.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	// TODO: make something to avoid infinite sessions for a user
	sessionId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(user.ID)
	if err == nil {
		return nil, err
	}

	session, err = createTokens(sessionId, userId)
	if err != nil {
		return nil, err
	}

	err = s.queries.CreateSession(s.ctx, database.CreateSessionParams{
		ID:           session.ID.String(),
		RefreshToken: session.RefreshToken,
		AccessToken:  session.AccessToken,
		UserID:       session.UserID.String(),
	})
	if err != nil {
		return nil, err
	}

	tokens = &types.AuthResponse{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	}

	return tokens, nil
}

func (s *authStore) SignUp(r *types.SignUpRequest) (*types.AuthResponse, error) {
	tokens := new(types.AuthResponse)
	session := new(types.Session)

	_, err := s.queries.GetUser(s.ctx, database.GetUserParams{
		Email:    r.Email,
		Username: r.Email,
	})
	if err == nil {
		return nil, errors.New("a user with this credentials already exists")
	}

	newUser, err := types.CreateUser(r)
	if err == nil {
		return nil, err
	}

	sessionId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(newUser.ID)
	if err == nil {
		return nil, err
	}

	session, err = createTokens(sessionId, userId)
	if err != nil {
		return nil, err
	}

	err = s.queries.CreateUser(s.ctx, *newUser)
	if err != nil {
		return nil, err
	}

	err = s.queries.CreateSession(s.ctx, database.CreateSessionParams{
		ID:           session.ID.String(),
		RefreshToken: session.RefreshToken,
		AccessToken:  session.AccessToken,
		UserID:       session.UserID.String(),
	})
	if err != nil {
		return nil, err
	}

	tokens = &types.AuthResponse{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	}

	return tokens, nil
}

func (s *authStore) Refresh(session *types.Session) (*types.AuthResponse, error) {
	tokens := new(types.AuthResponse)

	newSession, err := createTokens(session.ID, session.UserID)
	if err != nil {
		return nil, err
	}

	err = s.queries.UpdateSession(s.ctx, database.UpdateSessionParams{
		RefreshToken: newSession.RefreshToken,
		AccessToken:  newSession.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	tokens = &types.AuthResponse{
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
	}

	return tokens, nil
}

func (s *authStore) RecoverPassword(r *types.RecoverPasswordRequest) (msg string, err error) {
	return "not yet implemented", nil
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

func createTokens(sessionId, userId uuid.UUID) (*types.Session, error) {

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
		UserID:       userId,
	}, nil
}
