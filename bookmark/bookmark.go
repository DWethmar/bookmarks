package bookmark

import (
	"log/slog"
	"time"
)

// Bookmark is a struct that represents a bookmark.
type Bookmark struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

// Store is an interface that represents a bookmark store.
type Store interface {
	Add(b *Bookmark) error
	List() ([]*Bookmark, error)
	Delete(title string) error
}

// Library is a struct that represents a bookmark library.
type Library struct {
	logger *slog.Logger
	store  Store
}

func NewLibrary(logger *slog.Logger, store Store) *Library {
	return &Library{
		logger: logger,
		store:  store,
	}
}

// Add adds a bookmark to the library.
func (l *Library) Add(b *Bookmark) error {
	return l.store.Add(b)
}

// List lists all bookmarks in the library.
func (l *Library) List() ([]*Bookmark, error) {
	return l.store.List()
}

// Delete deletes a bookmark from the library.
func (l *Library) Delete(title string) error {
	return l.store.Delete(title)
}
