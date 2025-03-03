package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define the request and response structures
type DeepSeekRequest struct {
	Prompt string `json:"prompt"`
	// Add other required fields based on DeepSeek's API documentation
}

type DeepSeekResponse struct {
	Result string `json:"result"`
	// Add other fields based on DeepSeek's API documentation
}

func main() {

	// Replace with your actual API endpoint and key
	apiURL := "https://api.deepseek.com/v1/chat/completions"
	apiKey := "your-api-key-here"

	// Create the request payload and send it to the API
	requestBody := DeepSeekRequest{
		Prompt: "Hello, DeepSeek!",
	}

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling request body to JSON:", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set the required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return

	}

	// Unmarshal the response body to the DeepSeekResponse struct
	var response DeepSeekResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return
	}

	// Print the result
	fmt.Println("DeepSeek Response: ", response.Result)

}
