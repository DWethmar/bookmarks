package json_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/DWethmar/bookmarks/bookmark"
	"github.com/DWethmar/bookmarks/bookmark/json"
	"github.com/google/go-cmp/cmp"
)

func TestBookmark_Map(t *testing.T) {
	type fields struct {
		Title     string
		Content   string
		CreatedAt time.Time
	}
	type args struct {
		i *bookmark.Bookmark
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *json.Bookmark
	}{
		{
			name: "Test 1",
			fields: fields{
				Title:     "Test 1",
				Content:   "Test 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				i: &bookmark.Bookmark{
					Title:     "Test 1",
					Content:   "Test 1",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want: &json.Bookmark{
				Title:     "Test 1",
				Content:   "Test 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &json.Bookmark{
				Title:     tt.fields.Title,
				Content:   tt.fields.Content,
				CreatedAt: tt.fields.CreatedAt,
			}
			b.Map(tt.args.i)
			if diff := cmp.Diff(b, tt.want); diff != "" {
				t.Errorf("Bookmark.Map() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestBookmark_Unmap(t *testing.T) {
	type fields struct {
		Title     string
		Content   string
		CreatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *bookmark.Bookmark
	}{
		{
			name: "Test 1",
			fields: fields{
				Title:     "Test 1",
				Content:   "Test 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: &bookmark.Bookmark{
				Title:     "Test 1",
				Content:   "Test 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &json.Bookmark{
				Title:     tt.fields.Title,
				Content:   tt.fields.Content,
				CreatedAt: tt.fields.CreatedAt,
			}
			got := b.Unmap()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Bookmark.Unmap() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestStore_Add(t *testing.T) {
	t.Run("add should create a new bookmark", func(t *testing.T) {
		filePath := path.Join(t.TempDir(), "test.json")
		store := json.NewStore(filePath)
		b := &bookmark.Bookmark{
			Title:     "Test 1",
			Content:   "Test 1",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		if err := store.Add(b); err != nil {
			t.Errorf("Store.Add() error = %v", err)
		}
		expect := `[
  {
    "title": "Test 1",
    "content": "Test 1",
    "created_at": "2021-01-01T00:00:00Z"
  }
]
` // trailing newline
		file, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("failed to read file: %v", err)
		}

		if diff := cmp.Diff(string(file), expect); diff != "" {
			t.Errorf("Store.Add() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestStore_Delete(t *testing.T) {
	t.Run("add should delete a bookmark", func(t *testing.T) {
		filePath := path.Join(t.TempDir(), "test.json")
		store := json.NewStore(filePath)
		for i := range 3 {
			b := &bookmark.Bookmark{
				Title:     fmt.Sprintf("Test %d", i),
				Content:   fmt.Sprintf("Test %d", i),
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, i, time.UTC),
			}
			if err := store.Add(b); err != nil {
				t.Errorf("Store.Add() error = %v", err)
			}
		}
		if err := store.Delete("Test 1"); err != nil {
			t.Errorf("Store.Delete() error = %v", err)
		}
		expect := `[
  {
    "title": "Test 0",
    "content": "Test 0",
    "created_at": "2021-01-01T00:00:00Z"
  },
  {
    "title": "Test 2",
    "content": "Test 2",
    "created_at": "2021-01-01T00:00:00.000000002Z"
  }
]
` // trailing newline
		file, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("failed to read file: %v", err)
		}
		if diff := cmp.Diff(string(file), expect); diff != "" {
			t.Errorf("Store.Delete() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestStore_List(t *testing.T) {
	t.Run("add should list bookmarks", func(t *testing.T) {
		filePath := path.Join(t.TempDir(), "test.json")
		store := json.NewStore(filePath)
		for i := range 2 {
			b := &bookmark.Bookmark{
				Title:     fmt.Sprintf("Test %d", i),
				Content:   fmt.Sprintf("Test %d", i),
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, i, time.UTC),
			}
			if err := store.Add(b); err != nil {
				t.Errorf("Store.Add() error = %v", err)
			}
		}
		expect := `[
  {
    "title": "Test 0",
    "content": "Test 0",
    "created_at": "2021-01-01T00:00:00Z"
  },
  {
    "title": "Test 1",
    "content": "Test 1",
    "created_at": "2021-01-01T00:00:00.000000001Z"
  }
]
` // trailing newline
		file, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("failed to read file: %v", err)
		}
		if diff := cmp.Diff(string(file), expect); diff != "" {
			t.Errorf("Store.Delete() mismatch (-want +got):\n%s", diff)
		}
	})
}
