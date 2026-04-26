package main

import (
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	logger *slog.Logger
}

func main() {
	app := App{}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app.logger = logger

	server := http.Server{
		// If address is changed, ensure changes to playerForm url
		// port are also made. Currently not dynamic
		Addr:    ":3000",
		Handler: app.Routes(),
	}
	app.logger.Info("Starting server on port :3000")
	err := server.ListenAndServe()
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}

}
