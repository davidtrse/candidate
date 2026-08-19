package repository

import (
	"context"
	"errors"
	"time"

	"goassessment/internal/model"
)

// ErrNotFound is returned when no user matches the given id.
var ErrNotFound = errors.New("user not found")

var seed = map[int64]model.User{
	1: {ID: 1, Name: "Giap", Email: "giap@example.com"},
}

// GetUser looks up a user by id.
//
// TODO(candidate): this ignores ctx entirely, so a cancelled or timed-out
// request has no way to stop the (simulated) slow query below. Make it
// respect ctx cancellation/deadline and return ctx.Err() if it fires before
// the query finishes.
func GetUser(ctx context.Context, id int64) (model.User, error) {
	simulateSlowQuery()

	u, ok := seed[id]
	if !ok {
		return model.User{}, ErrNotFound
	}
	return u, nil
}

func simulateSlowQuery() {
	time.Sleep(1 * time.Second) // pretend this is a slow DB call
}
