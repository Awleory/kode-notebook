package v1

import (
	"context"
	"net/http"

	"github.com/awleory/kode/notebook/internal/entity"
	"github.com/gorilla/mux"
)

type User interface {
	CreateUser(ctx context.Context, inp entity.SignUpInput) error
	VerifyUser(ctx context.Context, email, pass string) (int, error)
}

type Note interface {
	CreateNote(ctx context.Context, inp entity.NoteCreating) error
	GetNotes(ctx context.Context, ownerId int) ([]entity.Note, error)
}

type Handler struct {
	usersService User
	notesService Note
}

func NewHandler(users User, notes Note) *Handler {
	return &Handler{
		usersService: users,
		notesService: notes,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(logging)

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	}

	note := router.PathPrefix("/note").Subrouter()
	note.Use(h.basicAuth)
	{
		note.HandleFunc("", h.createNote).Methods(http.MethodPost)
		note.HandleFunc("", h.getNotes).Methods(http.MethodGet)
	}

	return router
}
