package api

import (
	"fmt"
	"net/http"

	"github.com/hacdan/shellcheck_api/db"
	"github.com/hacdan/shellcheck_api/types"
)

func HandleCode(w http.ResponseWriter, r *http.Request) {
	codeString := r.PathValue("code")
	allCodes := db.ParseSCFile() //TODO: Move to interface/struct
	description, ok := allCodes[codeString]
	if !ok {
		respError(w, "Not found", http.StatusNotFound)
		return
	}

	code := types.SCCodeInfo{
		Code:        codeString,
		Description: description,
		Link:        fmt.Sprintf(SC_BASE_URL, codeString),
	}
	respondJSON(w, code, http.StatusOK)
}
