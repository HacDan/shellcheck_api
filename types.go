package main

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

