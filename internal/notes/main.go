package notes

import (
	"errors"

	"github.com/google/uuid"
)

type Registry interface {
	// Create create a new note from the given text returning its id.
	// Raises: ErrNoText, ErrTooLong.
	Create(text string) (string, error)
	// Delete deletes the note by id and returns its contents.
	// Raises: ErrInvalidID, ErrNotFound.
	Delete(id string) (string, error)
}

var (
	ErrNoText    = errors.New("no text")
	ErrTooLong   = errors.New("too long")
	ErrInvalidID = errors.New("invalid id")
	ErrNotFound  = errors.New("not found")
)

type RAMRegistry struct {
	MaxNoteLength uint64
	Notes         map[uuid.UUID]string
}

func NewRAMRegistry(maxNoteLength uint64) *RAMRegistry {
	if maxNoteLength < 1 {
		panic("invalid max note length")
	}

	return &RAMRegistry{
		MaxNoteLength: maxNoteLength,
		Notes:         make(map[uuid.UUID]string),
	}
}

func (reg *RAMRegistry) Create(text string) (string, error) {
	textLength := uint64(len(text))
	if textLength == 0 {
		return "", ErrNoText
	}
	if textLength > reg.MaxNoteLength {
		return "", ErrTooLong
	}

	id := uuid.New()
	reg.Notes[id] = text

	return id.String(), nil
}

func (reg *RAMRegistry) Delete(id string) (string, error) {
	id_, err := uuid.Parse(id)
	if err != nil {
		return "", ErrInvalidID
	}

	text, found := reg.Notes[id_]
	if !found {
		return "", ErrNotFound
	}
	delete(reg.Notes, id_)

	return text, nil
}
