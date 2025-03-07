package bookmark_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DWethmar/bookmarks/bookmark"
)

const htmlContent = `
<!DOCTYPE html>
<html>
<head>
	<title>Example Domain</title>
</head>
<body>
	<div>
		<h1>webpage</h1>
	</div>
</body>
</html>
`

func TestFetchTitle(t *testing.T) {
	t.Run("Successfully fetch webpage", func(t *testing.T) {
		// Create a test server that serves the mock HTML
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(htmlContent))
		}))
		defer server.Close()

		// Use the test server's URL instead of an actual URL
		client := &http.Client{}
		title, err := bookmark.FetchTitle(t.Context(), client, server.URL)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedTitle := "Example Domain"
		if title != expectedTitle {
			t.Errorf("expected title %q, got %q", expectedTitle, title)
		}
	})
}
