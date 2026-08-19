package service

import (
	"context"

	"goassessment/internal/model"
	"goassessment/internal/repository"
)

// GetUser fetches a user by id and lets the caller decide how to react to
// errors (including repository.ErrNotFound).
//
// TODO(candidate): the current version panics on error, which is wrong for
// an expected business case like "not found". Change the signature to
// return (model.User, error) instead of panicking, and propagate ctx and
// the repository error unchanged.
func GetUser(ctx context.Context, id int64) model.User {
	u, err := repository.GetUser(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}
