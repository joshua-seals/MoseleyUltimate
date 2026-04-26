package main

import (
	"net/http"
)

func (app *App) Routes() *http.ServeMux {
	r := http.NewServeMux()
	// Fileserver for assets
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	// General routes go here
	r.HandleFunc("GET /{$}", app.home)
	r.HandleFunc("GET /about", app.about)
	r.HandleFunc("GET /ultimate", app.ultimate)
	r.HandleFunc("GET /players", app.players)

	// r.HandleFunc("GET /login", app.login)
	// r.HandleFunc("POST /signup", app.signUp)

	r.HandleFunc("POST /ultimate/game/{status}", app.gameStatus)

	// Admin routes
	r.HandleFunc("GET /admin", app.adminDashboard)
	r.HandleFunc("GET /admin/login", app.adminLoginPage)
	r.HandleFunc("POST /admin/login", app.adminLogin)
	r.HandleFunc("POST /admin/logout", app.adminLogout)
	r.HandleFunc("POST /admin/game-status", app.adminGameStatus)

	// r.HandleFunc("GET /players/new", app.playerForm)
	// r.HandleFunc("POST /players/new", app.regsiterPlayerForm)

	// Only for use after creation of a new player calls the updated card.
	// r.HandleFunc("GET /players/{id}", app.playerCard)
	return r
}
