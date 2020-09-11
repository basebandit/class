// Package user contains user related CRUD functionality.
package user

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	// ErrNotFound is used when a specific User is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")
)

// Create inserts a new user into the database.
func Create(ctx context.Context, db *sqlx.DB, nu NewUser, now time.Time) (User, error) {

	return User{}, nil
}
