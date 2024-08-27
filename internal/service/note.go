package service

import (
	"context"

	"github.com/awleory/kode/notebook/internal/entity"
	"github.com/awleory/kode/notebook/pkg/yandex/speller"
)

type NoteRepo interface {
	CreateNote(ctx context.Context, note entity.NoteCreating) error
	GetNotes(ctx context.Context, ownerId int) ([]entity.Note, error)
}

type NoteService struct {
	repo NoteRepo
}

func NewNote(repo NoteRepo) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) CreateNote(ctx context.Context, note entity.NoteCreating) error {
	text, err := speller.CheckText(ctx, note.Text)
	if err != nil {
		return err
	}
	note.Text = text
	return s.repo.CreateNote(ctx, note)
}

func (s *NoteService) GetNotes(ctx context.Context, ownerId int) ([]entity.Note, error) {
	return s.repo.GetNotes(ctx, ownerId)
}
