package bookmark

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// FetchTitle retrieves the title of the given URL using an HTTP client.
func FetchTitle(ctx context.Context, client *http.Client, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "bookmarks") // Custom User-Agent

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.TextToken, html.EndTagToken, html.SelfClosingTagToken, html.CommentToken, html.DoctypeToken:
			continue
		case html.ErrorToken:
			return "", errors.New("title not found")
		case html.StartTagToken:
			tok := z.Token()
			if tok.Data == "title" {
				tt = z.Next()
				if tt == html.TextToken {
					return strings.TrimSpace(z.Token().Data), nil // Trim spaces from title
				}
			}
		}
	}
}
