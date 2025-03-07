package json

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/DWethmar/bookmarks/bookmark"
)

var _ bookmark.Store = &Store{}

var (
	// ErrNotFound is returned when a bookmark is not found.
	ErrNotFound = errors.New("bookmark not found")
)

// Bookmark is a struct that represents a bookmark.
type Bookmark struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Map maps a bookmark.Bookmark to a Bookmark.
func (b *Bookmark) Map(i *bookmark.Bookmark) {
	b.Title = i.Title
	b.Content = i.Content
	b.CreatedAt = i.CreatedAt
}

// Unmap maps a Bookmark to a bookmark.Bookmark.
func (b *Bookmark) Unmap() *bookmark.Bookmark {
	return &bookmark.Bookmark{
		Title:     b.Title,
		Content:   b.Content,
		CreatedAt: b.CreatedAt,
	}
}

// Store is a struct that represents a json store.
type Store struct {
	filePath string
	mutex    sync.Mutex
}

// NewStore creates a new json store.
func NewStore(filePath string) *Store {
	return &Store{
		filePath: filePath,
	}
}

// Add implements bookmark.Store.
func (s *Store) Add(b *bookmark.Bookmark) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	e := &Bookmark{}
	e.Map(b)

	// Load existing bookmarks
	bookmarks, err := s.load()
	if err != nil {
		return err
	}

	// Append the new bookmark
	bookmarks = append(bookmarks, e)

	// Save back to file
	return s.save(bookmarks)
}

// Delete implements bookmark.Store.
func (s *Store) Delete(title string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// Load existing bookmarks
	bookmarks, err := s.load()
	if err != nil {
		return err
	}
	// Filter out the bookmark with the given title
	var newBookmarks []*Bookmark
	for _, b := range bookmarks {
		if b.Title != title {
			newBookmarks = append(newBookmarks, b)
		}
	}
	if len(newBookmarks) == len(bookmarks) {
		return ErrNotFound
	}
	// Save back to file
	return s.save(newBookmarks)
}

// List implements bookmark.Store.
func (s *Store) List() ([]*bookmark.Bookmark, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// Load existing bookmarks
	r, err := s.load()
	if err != nil {
		return nil, err
	}
	// Map the bookmarks
	var bookmarks []*bookmark.Bookmark
	for _, b := range r {
		bookmarks = append(bookmarks, b.Unmap())
	}
	return bookmarks, nil
}

// load reads the JSON file and returns the list of bookmarks.
func (s *Store) load() ([]*Bookmark, error) {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []*Bookmark{}, nil
	}
	file, err := os.Open(s.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var bookmarks []*Bookmark
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&bookmarks); err != nil {
		return nil, err
	}
	return bookmarks, nil
}

// save writes the list of bookmarks to the JSON file.
func (s *Store) save(bookmarks []*Bookmark) error {
	file, err := os.Create(s.filePath) // Overwrite the file
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	return encoder.Encode(bookmarks)
}
