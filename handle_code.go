package main

import (
	"fmt"
	"net/http"
)

func handleCode(w http.ResponseWriter, r *http.Request) {
	codeString := r.PathValue("code")
	allCodes := parseSCFile() //TODO: Move to interface/struct
	description, ok := allCodes[codeString]
	if !ok {
		respError(w, "Not found", http.StatusNotFound)
		return
	}

	code := SCCodeInfo{
		Code:        codeString,
		Description: description,
		Link:        fmt.Sprintf(SC_BASE_URL, codeString),
	}
	respondJSON(w, code, http.StatusOK)
}
