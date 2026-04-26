package main

import (
	"github.com/joshua-seals/MoseleyUltimate/internal/models"
)

// Add all templateData structs to this struct
// This way it is all accessible.
type templateData struct {
	Player     *models.Player
	Players    []*models.Player
	GameStatus string
	Date       string
	Flash      string
}

// Use later once front-end work is complete
//
// func newTemplateCache() (map[string]*template.Template, error) {
// 	cache := map[string]*template.Template{}

// 	pages, err := filepath.Glob("./ui/html/pages/*.html")
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, page := range pages {
// 		name := filepath.Base(page)
// 		files := []string{
// 			"./ui/html/base.html",
// 			"./ui/html/partials.html",
// 			page,
// 		}

// 		ts, err := template.ParseFiles(files...)
// 		if err != nil {
// 			return nil, err
// 		}
// 		cache[name] = ts
// 	}
// 	return cache, nil
// }
