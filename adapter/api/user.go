package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
