package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/usecase"
)

type PostApi struct {
	uc usecase.PostUseCase
}

func NewPostApi(uc usecase.PostUseCase) *PostApi {
	return &PostApi{
		uc: uc,
	}
}

func (a *PostApi) FetchByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid user ID", "id", chi.URLParam(r, "user_id"), "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	posts, err := a.uc.FetchByUserID(r.Context(), userID)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to fetch posts", "user_id", userID, "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToFetch)
		return
	}
	if posts == nil {
		posts = []*models.Post{}
	}
	WriteJSON(w, http.StatusOK, posts)
}

func (a *PostApi) FetchByUserIDWithComments(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid user ID", "id", chi.URLParam(r, "user_id"), "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	posts, err := a.uc.FetchByUserIDWithComments(r.Context(), userID)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to fetch posts with comments", "user_id", userID, "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToFetch)
		return
	}
	if posts == nil {
		posts = []*models.PostWithComments{}
	}
	WriteJSON(w, http.StatusOK, posts)
}
