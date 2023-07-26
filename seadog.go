package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	apiKey := ""

	// Parse command-line arguments
	cmd := flag.String("cmd", "", "Command to execute")
	dataSource := flag.String("url", "", "Source URL")
	dataType := flag.String("type", "", "ingestion type")
	flag.Parse()

	// Execute the specified command
	switch *cmd {
	case "start-conversation":
		response, err := newConversation(apiKey)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		printResponse(response)
		// conversationID := response["conversation_id"].(string)

	// case "chat":
	// 	// get the conversation ID

	// 	// get the question from the command line
	// 	question := flag.Arg(1)
	// 	if question == "" {
	// 		fmt.Fprintln(os.Stderr, "question is required")
	// 		os.Exit(1)
	// 	}

	// 	chatResponse, err := chat(apiKey, question, question , conversationID)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	// 	printResponse(chatResponse)

	case "ingest":
		if *dataSource == "" || *dataType == "" {
			fmt.Fprintln(os.Stderr, "url and Type are required")
			os.Exit(1)
		}
		response, err := ingestData(apiKey, *dataSource, *dataType)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		printResponse(response)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		os.Exit(1)
	}
}

func newConversation(apiKey string) (map[string]interface{}, error) {
	url := "https://api.mendable.ai/v0/newConversation"

	// Create the request body
	requestBody := map[string]interface{}{
		"api_key": apiKey,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Send the HTTP request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func ingestData(apiKey string, dataSource string, dataType string) (map[string]interface{}, error) {
	url := "https://api.mendable.ai/v0/ingestData"

	// Create the request body
	requestBody := map[string]interface{}{
		"api_key": apiKey,
		"url":     dataSource,
		"type":    dataType,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Send the HTTP request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func chat(apiKey string, question string, history []map[string]string, conversationID int) (string, error) {
	url := "https://api.mendable.ai/v0/mendableChat"

	data := map[string]interface{}{
		"api_key":         apiKey,
		"question":        question,
		"history":         history,
		"conversation_id": conversationID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func printResponse(response map[string]interface{}) {
	// Convert the response map into a JSON string and print the response
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(responseJSON))
}

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	apiKey := "d67542ec-87c7-489e-8c66-24895eb57136"

// 	// Send a newConversation request
// 	newConversationRequestBody := map[string]interface{}{
// 		"api_key": apiKey,
// 	}
// 	newConversationResponse, err := newConversation(apiKey, "newConversation", newConversationRequestBody)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// convert the response map into a JSON string and print the response
// 	newConversationResponseJSON, err := json.Marshal(newConversationResponse)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(newConversationResponseJSON))

// 	// Now let's ingest some data
// 	ingestResponse, err := ingestData(apiKey, "https://api.mendable.ai/v0/ingestData", "url")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// convert the response map into a JSON string and print the response
// 	ingestResponseJSON, err := json.Marshal(ingestResponse)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(ingestResponseJSON))

// 	// Get the task ID from the ingest response
// 	taskID := ingestResponse["task_id"].(string)

// 	// Now let's check the status of the task
// 	ingestErr := fetchIngestionStatus(taskID)
// 	if ingestErr != nil {
// 		fmt.Println(ingestErr)
// 	}

// }

// func newConversation(apiKey string, endpoint string, requestBody map[string]interface{}) (map[string]interface{}, error) {
// 	url := fmt.Sprintf("https://api.mendable.ai/v0/%s", endpoint)

// 	// Create the request body
// 	requestBodyBytes, err := json.Marshal(requestBody)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Send the HTTP request
// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	var response map[string]interface{}
// 	err = json.NewDecoder(resp.Body).Decode(&response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }

// func ingestData(apiKey string, url string, ingestionType string) (map[string]interface{}, error) {
// 	// Create the request body
// 	requestBody := map[string]string{
// 		"api_key": apiKey,
// 		"url":     url,
// 		"type":    ingestionType,
// 	}
// 	requestBodyBytes, err := json.Marshal(requestBody)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Send the HTTP request
// 	resp, err := http.Post("https://api.mendable.ai/v0/ingestData", "application/json", bytes.NewBuffer(requestBodyBytes))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	var response map[string]interface{}
// 	err = json.NewDecoder(resp.Body).Decode(&response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }

// func fetchIngestionStatus(taskID string) error {
// 	url := "https://api.mendable.ai/v0/ingestionStatus"

// 	data := map[string]string{
// 		"task_id": taskID,
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return fmt.Errorf("error marshalling JSON: %v", err)
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return fmt.Errorf("error creating request: %v", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("error sending request: %v", err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("error reading response body: %v", err)
// 	}

// 	fmt.Println(string(body))

// 	return nil
// }
