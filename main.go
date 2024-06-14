package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

const (
	SC_BASE_URL             = "https://github.com/koalaman/shellcheck/wiki/%s"
	ContentTypeJSON  = "application/json"
	DefaultListenPort= ":8888"
	ErrorParsingJSON = "Error parsing JSON"
	ErrorNotFound    = "Not found"
)

type Err struct {
	Error string `json:"error"`
}

type Parsecode struct {
	Codes string `json:"codes"`
}

type SCCodeInfo struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Link        string `json:"link"`
}


func main() {
	http.HandleFunc("/api/v1/codes", handleAllCodes)
	http.HandleFunc("/api/v1/codes/{code}", handleCode)
	http.HandleFunc("/api/v1/codes/parse", handleParse)
	log.Fatal(http.ListenAndServe(":8888", nil)) //TODO: Change to environment variable/.env with a default
}

func handleCode(w http.ResponseWriter, r *http.Request) {
	codeString := r.PathValue("code")
	allCodes := parseSCFile() //TODO: Move to interface/struct
	description, ok := allCodes[codeString]
	if !ok {
		respError(w, "Code not found", http.StatusNotFound)
		return
	}

	code := SCCodeInfo{
		Code:        codeString,
		Description: description,
		Link:        fmt.Sprintf(SC_BASE_URL, codeString),
	}
	respondJSON(w, code, http.StatusOK)
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

	respondJSON(w, codes, http.StatusOK)
}

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
	mErr := Err{
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
