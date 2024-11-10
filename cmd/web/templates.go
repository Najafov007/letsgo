package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.nijat.net/internal/models"
	"time"
)

type templateData struct{
	Snippet 	models.Snippet
	Snippets 	[]models.Snippet
	CurrentYear int
	Form 		any
	Flash 		string
	IsAuthenticated	bool
	CSRFToken	string
}

func humanDate(t time.Time) string{
	return t.Format("02 Jan 2006 at 15:06")
} 

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
		if err != nil {
			return nil, err
	
	}

	for _, page := range pages {
		name := filepath.Base(page)

	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/partials/navigation.html",
	// 	page,
	// }

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
