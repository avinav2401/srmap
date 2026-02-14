package main

import (
	"goscraper/src/bootstrap"
	"net/http"

	"github.com/gofiber/adaptor/v2"
)

// Handler is the entry point for Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	app := bootstrap.SetupApp()
	http.StripPrefix("/api", adaptor.FiberApp(app)).ServeHTTP(w, r)
}
