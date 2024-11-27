package mock

import (
	"time"

	"github.com/jfilipedias/snippetbox/internal/model"
)

var mockSnippet = model.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (model.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return model.Snippet{}, model.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]model.Snippet, error) {
	return []model.Snippet{mockSnippet}, nil
}
