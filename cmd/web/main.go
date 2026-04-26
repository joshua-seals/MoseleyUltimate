package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joshua-seals/MoseleyUltimate/internal/models"
)

type App struct {
	logger *slog.Logger
}

func main() {
	app := App{}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app.logger = logger

	adminPass := os.Getenv("ADMIN_PASSWORD")
	if adminPass == "" {
		adminPass = "changeme"
		logger.Warn("ADMIN_PASSWORD not set, using default. Set ADMIN_PASSWORD env var in production.")
	}
	if err := models.InitAdmin(adminPass); err != nil {
		logger.Error("Failed to initialize admin", "error", err)
		os.Exit(1)
	}

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
