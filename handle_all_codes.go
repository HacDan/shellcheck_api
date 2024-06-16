package main

import (
	"fmt"
	"net/http"
	"sort"
)


func handleAllCodes(w http.ResponseWriter, r *http.Request) {
	var codes []SCCodeInfo
	allCodes := parseSCFile()

	for code, description := range allCodes {
		sccodeinfo := SCCodeInfo{
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
