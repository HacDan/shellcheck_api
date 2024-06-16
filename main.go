package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

const (
	SC_BASE_URL       = "https://github.com/koalaman/shellcheck/wiki/%s"
	DefaultListenPort = ":8080"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/v1/codes", handleAllCodes)
	http.HandleFunc("/api/v1/codes/{code}", handleCode)
	http.HandleFunc("/api/v1/codes/parse", handleParse)

	log.Fatal(http.ListenAndServe(getEnv("ADDRESS", DefaultListenPort), nil))
}
