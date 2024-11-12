package main

import "github.com/jfilipedias/snippetbox/internal/model"

type templateData struct {
	Snippet  model.Snippet
	Snippets []model.Snippet
}
