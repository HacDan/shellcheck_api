package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

const SC_BASE_URL = "https://github.com/koalaman/shellcheck/wiki/%s"

type SCCodeInfo struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Parsecode struct {
	Codes string `json:"codes"`
}

type Err struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/api/v1/codes", handleAllCodes)
	http.HandleFunc("/api/v1/codes/{code}", handleCode)
	http.HandleFunc("/api/v1/codes/parse", handleParse)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func handleCode(w http.ResponseWriter, r *http.Request) {
	codeString := r.PathValue("code")
	allCodes := parseSCFile() //TODO: Move to interface/struct
	if _, ok := allCodes[codeString]; !ok {
		respError(w, "Not found")
		return
	}

	code := SCCodeInfo{
		Code:        codeString,
		Description: allCodes[codeString],
		Link:        fmt.Sprintf(SC_BASE_URL, codeString),
	}

	jsonCode, err := json.Marshal(code)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCode)

}

func respError(w http.ResponseWriter, errorString string) {
	mErr := Err{
		Error: errorString,
	}
	errorJson, err := json.Marshal(mErr)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(errorJson)
}

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

	jsonCodes, err := json.Marshal(codes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCodes)

}

func handleParse(w http.ResponseWriter, r *http.Request) {
	var parseCode Parsecode
	var scCodeInfos []SCCodeInfo

	err := json.NewDecoder(r.Body).Decode(&parseCode)
	if err != nil {
		respError(w, "Error parsing json")
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

	jsonCodes, err := json.Marshal(scCodeInfos)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCodes)
}
