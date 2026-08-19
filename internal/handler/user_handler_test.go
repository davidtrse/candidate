package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"goassessment/internal/handler"
)

func TestGetUserHandler_InvalidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users?id=abc", nil)
	rec := httptest.NewRecorder()

	handler.GetUserHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid id, got %d", rec.Code)
	}
}

func TestGetUserHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users?id=999", nil)
	rec := httptest.NewRecorder()

	handler.GetUserHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for missing user, got %d", rec.Code)
	}
}

func TestGetUserHandler_Found(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users?id=1", nil)
	rec := httptest.NewRecorder()

	handler.GetUserHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 for existing user, got %d", rec.Code)
	}
}
