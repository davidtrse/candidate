package service_test

import (
	"context"
	"errors"
	"testing"

	"goassessment/internal/repository"
	"goassessment/internal/service"
)

func TestGetUser_NotFoundReturnsError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("GetUser should return an error for a missing user, not panic: %v", r)
		}
	}()

	_, err := service.GetUser(context.Background(), 999)
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("expected repository.ErrNotFound, got %v", err)
	}
}
