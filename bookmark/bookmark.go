package bookmark

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
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
	client *http.Client
}

func NewLibrary(logger *slog.Logger, store Store) *Library {
	return &Library{
		logger: logger,
		store:  store,
		client: &http.Client{},
	}
}

// Add adds a bookmark to the library.
func (l *Library) Add(ctx context.Context, b *Bookmark) error {
	if b.Title == "" && isURL(b.Content) {
		title, err := FetchTitle(ctx, l.client, b.Content)
		if err != nil {
			return err
		}
		b.Title = title
	}
	return l.store.Add(b)
}

// List lists all bookmarks in the library.
func (l *Library) List() ([]*Bookmark, error) {
	return l.store.List()
}

// Search searches for bookmarks in the library.
func (l *Library) Search(query string) ([]*Bookmark, error) {
	bookmarks, err := l.store.List()
	if err != nil {
		return nil, err
	}
	var results []*Bookmark
	for _, b := range bookmarks {
		if strings.Contains(b.Title, query) || strings.Contains(b.Content, query) {
			results = append(results, b)
		}
	}
	return results, nil
}

// Delete deletes a bookmark from the library.
func (l *Library) Delete(title string) error {
	return l.store.Delete(title)
}

// isURL checks if a string is a URL.
func isURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}
