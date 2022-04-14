package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/hafizmfadli/snippetbox/pkg/forms"
	"github.com/hafizmfadli/snippetbox/pkg/models"
)

// Define a templateData typoe to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	AuthenticatedUser *models.User
	CSRFToken string
	CurrentYear int
	Flash string
	Form *forms.Form
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func humanDate(t time.Time) string {
	// Return the empty string if time has the zero value.
	if t.IsZero() {
		return ""
	}

	// Convert the time to UTC before formatting it.
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		// extrac the file name (like 'home.page.tmpl') from the full path
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before call the ParseFiles()
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}