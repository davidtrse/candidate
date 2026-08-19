package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"goassessment/internal/repository"
	"goassessment/internal/service"
)

// GetUserHandler handles GET /users?id=<id>.
//
// TODO(candidate):
//  1. Validate the "id" query param — it must be a positive integer.
//     Return 400 Bad Request if it isn't.
//  2. Give the request a timeout (e.g. 2s) via context instead of using
//     context.Background() directly, so a slow downstream call can't hang
//     the handler forever.
//  3. Once service.GetUser returns (model.User, error) (see the TODO in
//     internal/service), map repository.ErrNotFound to 404 Not Found —
//     don't let it fall through as a 500.
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if id < 1 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2 * time.Second)
	defer cancel()

	u,e := service.GetUser(ctx, id)
	if e == repository.ErrNotFound {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
