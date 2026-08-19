package repository

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestGetUser_RespectsContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := GetUser(ctx, 1)
	elapsed := time.Since(start)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded, got %v", err)
	}
	if elapsed > 500*time.Millisecond {
		t.Fatalf("GetUser should have stopped when ctx timed out, took %v", elapsed)
	}
}

func TestGetUser_NotFound(t *testing.T) {
	ctx := context.Background()
	_, err := GetUser(ctx, 999)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestGetUser_Found(t *testing.T) {
	ctx := context.Background()
	u, err := GetUser(ctx, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Name != "Giap" {
		t.Fatalf("expected Name %q, got %q", "Giap", u.Name)
	}
}
