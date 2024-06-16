package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func handleParse(w http.ResponseWriter, r *http.Request) {
	var parseCode Parsecode
	var scCodeInfos []SCCodeInfo

	err := json.NewDecoder(r.Body).Decode(&parseCode)
	if err != nil {
		respError(w, "Error parsing json", http.StatusUnprocessableEntity)
		return
	}

	scCodes := findAllCodes(parseCode.Codes)
	allScCodes := parseSCFile()

	for _, scCode := range scCodes {
		if _, ok := allScCodes[scCode]; ok {
			tempCodeInfo := SCCodeInfo{
				Code:        scCode,
				Description: allScCodes[scCode],
				Link:        fmt.Sprintf(SC_BASE_URL, scCode),
			}
			scCodeInfos = append(scCodeInfos, tempCodeInfo)
		}
	}

	sort.Slice(scCodeInfos, func(i, j int) bool {
		return scCodeInfos[i].Code < scCodeInfos[j].Code
	})

	respondJSON(w, scCodeInfos, http.StatusOK)
}

