package data

import (
	"context"
	"database/sql"
	"gps-tracker/internal/database"
	"gps-tracker/internal/types"

	"github.com/google/uuid"
)

type UserStore interface {
	CheckUsedEmail(email string) (used bool, err error)
	CheckUsedUsername(username string) (used bool, err error)
	CheckUsedPhoneNumber(phoneNumber string) (used bool, err error)
	GetUser(id uuid.UUID) (user *types.User, err error)
	UpdateUser(r *types.UpdateUser) (err error)
	DeleteUser(id uuid.UUID) (err error)
}

func (s *storage) UserStore() UserStore {
	return NewUserStore(s.ctx, s.queries)
}

type userStore struct {
	ctx     context.Context
	queries *database.Queries
}

func NewUserStore(ctx context.Context, queries *database.Queries) UserStore {
	return &userStore{
		ctx:     ctx,
		queries: queries,
	}
}

func (s *userStore) CheckUsedEmail(email string) (used bool, err error) {
	_, err = s.queries.GetEmail(s.ctx, email)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return true, err
	}

	used = true

	return
}
func (s *userStore) CheckUsedUsername(username string) (used bool, err error) {
	_, err = s.queries.GetUsername(s.ctx, username)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return true, err
	}

	used = true

	return
}
func (s *userStore) CheckUsedPhoneNumber(phoneNumber string) (used bool, err error) {
	_, err = s.queries.GetPhoneNumber(s.ctx, phoneNumber)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return true, err
	}

	used = true

	return
}

func (s *userStore) GetUser(id uuid.UUID) (*types.User, error) {
	user := new(types.User)

	dbUser, err := s.queries.GetUserById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	user, err = types.MapDbUser(&dbUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userStore) UpdateUser(r *types.UpdateUser) (err error) {
	err = s.queries.UpdateUser(s.ctx, database.UpdateUserParams{
		ID:        r.ID.String(),
		Firstname: r.FirstName,
		Lastname:  r.LastName,
		Alias: sql.NullString{
			String: r.Alias,
			Valid:  r.Alias != "",
		},
		BirthDate: r.BirthDate,
	})
	if err != nil {
		return err
	}

	return
}
func (s *userStore) DeleteUser(id uuid.UUID) (err error) {
	err = s.queries.DeleteUser(s.ctx, id.String())
	if err != nil {
		return err
	}

	return
}
