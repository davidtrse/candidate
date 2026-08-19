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
	type Result struct {
		u model.User
		e error
	}

	buffered := make(chan Result, 1)

	go func() {
		simulateSlowQuery()
		u, ok := seed[id]
		if !ok {
			buffered <- Result{model.User{}, ErrNotFound}
			return
		}
		buffered <- Result{u, nil}
	}()

	select {

	case r := <- buffered:
		return r.u, r.e

	case <- ctx.Done():
		return model.User{}, ctx.Err()
	}
}

func simulateSlowQuery() {
	time.Sleep(1 * time.Second) // pretend this is a slow DB call
}
