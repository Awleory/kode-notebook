package psql

import (
	"context"
	"database/sql"

	"github.com/awleory/kode/notebook/internal/entity"
)

type Notes struct {
	db *sql.DB
}

func NewNotes(db *sql.DB) *Notes {
	return &Notes{
		db: db,
	}
}

func (r *Notes) CreateNote(ctx context.Context, note entity.NoteCreating) error {
	_, err := r.db.Exec("INSERT INTO notes (owner_id, title, text) values ($1, $2, $3)",
		note.OwnerId, note.Title, note.Text)

	return err
}

func (r *Notes) GetNotes(ctx context.Context, ownerId int) ([]entity.Note, error) {
	rows, err := r.db.Query("SELECT title, text FROM notes WHERE owner_id=$1", ownerId)
	if err != nil {
		return nil, err
	}

	notes := make([]entity.Note, 0)
	for rows.Next() {
		note := entity.Note{}
		rows.Scan(&note.Title, &note.Text)
		notes = append(notes, note)
	}

	return notes, err
}
