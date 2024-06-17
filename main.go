package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/hacdan/shellcheck_api/api"
	"github.com/hacdan/shellcheck_api/utils"
	"github.com/joho/godotenv"
)

const (
	DefaultListenPort = ":8080"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading environment variables", err)
	}

	http.HandleFunc("/api/v1/codes", api.HandleAllCodes)
	http.HandleFunc("/api/v1/codes/{code}", api.HandleCode)
	http.HandleFunc("/api/v1/codes/parse", api.HandleParse)

	slog.Info("Starting Server on localhost" + utils.GetEnv("ADDRESS", DefaultListenPort))
	slog.Error("Server failed: ", http.ListenAndServe(utils.GetEnv("ADDRESS", DefaultListenPort), nil))
	os.Exit(1)
}
