package main

import (
	"log/slog"
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
	server := api.NewServer(utils.GetEnv("ADDRESS", DefaultListenPort))

	slog.Info("Starting Server on localhost" + utils.GetEnv("ADDRESS", DefaultListenPort))
	slog.Error("Server failed: ", server.Start())
	os.Exit(1)
}
