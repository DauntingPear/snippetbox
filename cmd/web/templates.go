package main

import (
    "snippetbox.quackden.net/internal/models"
)

type templateData struct {
    Snippet models.Snippet
    Snippets []models.Snippet
}
