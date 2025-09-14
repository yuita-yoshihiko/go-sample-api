package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/usecase"
)

type UserApi struct {
	uc usecase.UserUseCase
}

func NewUserApi(uc usecase.UserUseCase) *UserApi {
	return &UserApi{uc: uc}
}

func (a *UserApi) Fetch(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid user ID", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}
	u, err := a.uc.Fetch(r.Context(), id)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to fetch user", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToFetch)
		return
	}
	WriteJSON(w, http.StatusOK, u)
}

func (a *UserApi) FetchWithPosts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid user ID", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}
	u, err := a.uc.FetchWithPosts(r.Context(), id)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to fetch user with posts", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToFetch)
		return
	}
	WriteJSON(w, http.StatusOK, u)
}

func (a *UserApi) Create(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.ErrorContext(r.Context(), "Invalid request body", "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}
	if err := a.uc.Create(r.Context(), user); err != nil {
		slog.ErrorContext(r.Context(), "Failed to create user", "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToPost)
		return
	}
	WriteJSON(w, http.StatusCreated, user)
}

func (a *UserApi) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid user ID", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}
	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.ErrorContext(r.Context(), "Invalid request body", "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}
	user.ID = id
	if err := a.uc.Update(r.Context(), user); err != nil {
		slog.ErrorContext(r.Context(), "Failed to update user", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToPost)
		return
	}
	WriteJSON(w, http.StatusOK, user)
}

func (a *UserApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid user ID", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}
	if err := a.uc.Delete(r.Context(), id); err != nil {
		slog.ErrorContext(r.Context(), "Failed to delete user", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToPost)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
