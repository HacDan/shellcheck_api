package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const SC_BASE_URL = "https://github.com/koalaman/shellcheck/wiki/%s"

type SCCodeInfo struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Err struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/api/v1/codes", handleAllCodes)
	http.HandleFunc("/api/v1/codes/{code}", handleCode)
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

	jsonCodes, err := json.Marshal(codes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCodes)

}
