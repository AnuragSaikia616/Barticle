package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// url of the api
const API_URL = "http://localhost:8000/"

// summary request
type SummaryInput struct {
	Text string `json:"text"`
}

// summary response
type SummaryResponse struct {
	Summary string `json:"summary"`
}

// summary client
func getSummary(text string) (string, error) {
	inputdata := SummaryInput{Text: text}
	jsonInput, err := json.Marshal(inputdata)
	if err != nil {
		return "", fmt.Errorf("Error marshalling input data", err)
	}
	resp, err := http.Post(API_URL+"summarize/", "application/json", bytes.NewBuffer(jsonInput))
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

// question answering request
type qaInput struct {
	Question string `json:"question"`
	Context  string `json:"context"`
}

// question answering response
type qaResponse struct {
	Answer string  `json:"answer"`
	Score  float32 `json:"score"`
	Start  int     `json:"start"`
	End    int     `json:"end"`
}

// question answreing client
func getAnswer(question string, context string) (string, error) {
	qaInput := qaInput{
		Question: question,
		Context:  context,
	}
	jsonInput, err := json.Marshal(qaInput)
	if err != nil {
		return "", fmt.Errorf("Error marshalling input data", err)
	}
	resp, err := http.Post(API_URL+"answer/", "application/json", bytes.NewBuffer(jsonInput))
	if err != nil {
		return "", fmt.Errorf("Error making POST request", err)
	}
	defer resp.Body.Close()

	var respAnswer qaResponse
	if err := json.NewDecoder(resp.Body).Decode(&respAnswer); err != nil {
		return "ERROR", fmt.Errorf("Error decoding the response body", err)
	}
	return respAnswer.Answer, nil
}

// scraper request
type scraperRequest struct {
	Url string `json:"url"`
}

// scraper response
type scraperResponse struct {
	Content string `json:"content"`
}

// scraper client
func getScrapedURLContent(url string) (string, error) {
	scraperRequest := scraperRequest{
		Url: url,
	}
	jsonInput, err := json.Marshal(scraperRequest)
	if err != nil {
		return "", fmt.Errorf("Error marshalling input data", err)
	}
	resp, err := http.Post(API_URL+"scrape/", "application/json", bytes.NewBuffer(jsonInput))
	if err != nil {
		return "", fmt.Errorf("Error making POST request", err)
	}
	defer resp.Body.Close()

	var respAnswer scraperResponse
	if err := json.NewDecoder(resp.Body).Decode(&respAnswer); err != nil {
		return "ERROR", fmt.Errorf("Error decoding the response body", err)
	}
	return respAnswer.Content, nil
}
