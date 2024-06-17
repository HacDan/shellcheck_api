package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/hacdan/shellcheck_api/db"
	"github.com/hacdan/shellcheck_api/types"
)

func HandleParse(w http.ResponseWriter, r *http.Request) {
	var parseCode types.Parsecode
	var scCodeInfos []types.SCCodeInfo

	err := json.NewDecoder(r.Body).Decode(&parseCode)
	if err != nil {
		respError(w, "Error parsing json", http.StatusUnprocessableEntity)
		return
	}

	scCodes := db.FindAllCodes(parseCode.Codes)
	allScCodes := db.ParseSCFile()

	for _, scCode := range scCodes {
		if _, ok := allScCodes[scCode]; ok {
			tempCodeInfo := types.SCCodeInfo{
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
