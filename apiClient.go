package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SummaryInput struct {
	Text string `json:"text"`
}
type SummaryResponse struct {
	Summary string `json:"summary"`
}
type qaInput struct {
	Question string `json:"question"`
	Context  string `json:"context"`
}
type qaResponse struct {
	Answer string `json:"answer"`
}

func getSummary(text string) (string, error) {
	url := "http://localhost:8000/summarize/"
	inputdata := SummaryInput{Text: text}
	jsonInput, err := json.Marshal(inputdata)
	if err != nil {
		return "", fmt.Errorf("Error marshalling input data", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonInput))
	if err != nil {
		return "", fmt.Errorf("Error making POST request", err)
	}
	defer resp.Body.Close()

	var respSummary SummaryResponse
	if err := json.NewDecoder(resp.Body).Decode(&respSummary); err != nil {
		return "", fmt.Errorf("Error decoding the response body", err)
	}
	return respSummary.Summary, nil
}
