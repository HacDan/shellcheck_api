package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	SC_BASE_URL       = "https://github.com/koalaman/shellcheck/wiki/%s"
	DefaultListenPort = ":8080"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading environment variables", err)
	}

	http.HandleFunc("/api/v1/codes", handleAllCodes)
	http.HandleFunc("/api/v1/codes/{code}", handleCode)
	http.HandleFunc("/api/v1/codes/parse", handleParse)

	slog.Error("Server failed: ", http.ListenAndServe(getEnv("ADDRESS", DefaultListenPort), nil))
	os.Exit(1)
}
