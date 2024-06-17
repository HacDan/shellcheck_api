package api

import (
	"encoding/json"
	"net/http"

	"github.com/hacdan/shellcheck_api/types"
)

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		respError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

func respError(w http.ResponseWriter, errorString string, status int) {
	mErr := types.Err{
		Error: errorString,
	}
	errorJson, err := json.Marshal(mErr)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(errorJson)
}
