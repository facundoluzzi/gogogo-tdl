package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TranslateCommand struct {
	Name string
}

const (
	apiURL             = "https://api.cognitive.microsofttranslator.com/translate?api-version=3.0&from=en&to=es"
	subscriptionKey    = "06e804bb63294eed9ccf2c8a7796c2d6"
	subscriptionRegion = "brazilsouth"
)

func (c *TranslateCommand) Run() error {
	file, err := os.Create(c.Name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	translateText := string(contentBytes)

	payload := []map[string]string{
		{"Text": translateText},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling payload: %v\n", err)
		os.Exit(1)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	// Set the required headers
	req.Header.Set("Ocp-Apim-Subscription-Key", subscriptionKey)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Ocp-Apim-Subscription-Region", subscriptionRegion)

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}
	err = os.WriteFile(c.Name, respBody, 0644)
	if err != nil {
		return fmt.Errorf("failed to write new content to file: %w", err)
	}
	return nil
}

func (c *TranslateCommand) Print() {
	print(c.Name)
}
