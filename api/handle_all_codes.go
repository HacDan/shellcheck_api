package api

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/hacdan/shellcheck_api/storage"
	"github.com/hacdan/shellcheck_api/types"
)

const SC_BASE_URL = "https://github.com/koalaman/shellcheck/wiki/%s"

func (s *Server) HandleAllCodes(w http.ResponseWriter, r *http.Request) {
	var codes []types.SCCodeInfo
	allCodes := storage.ParseSCFile()

	for code, description := range allCodes {
		sccodeinfo := types.SCCodeInfo{
			Code:        code,
			Description: description,
			Link:        fmt.Sprintf(SC_BASE_URL, code),
		}
		codes = append(codes, sccodeinfo)
	}

	sort.Slice(codes, func(i, j int) bool {
		return codes[i].Code < codes[j].Code
	})

	respondJSON(w, codes, http.StatusOK)
}
