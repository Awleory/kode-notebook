package v1

import (
	"encoding/json"
	"net/http"

	"github.com/awleory/kode/notebook/internal/entity"
)

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	userId, err := getCtxUserId(r.Context())
	if err != nil {
		logError("createNote", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var note entity.NoteCreating
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		logError("createNote", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	note.OwnerId = userId
	err = h.notesService.CreateNote(r.Context(), note)
	if err != nil {
		logError("createNote", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getNotes(w http.ResponseWriter, r *http.Request) {
	userId, err := getCtxUserId(r.Context())
	if err != nil {
		logError("getNotes", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	notes, err := h.notesService.GetNotes(r.Context(), userId)
	if err != nil {
		logError("getNotes", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(notes)
	if err != nil {
		logError("getNotes", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
