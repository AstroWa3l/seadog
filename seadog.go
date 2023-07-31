package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	apiKey := ""

	// Parse command-line arguments
	cmd := flag.String("cmd", "", "Command to execute")
	flag.Parse()

	// Execute the specified command
	switch *cmd {
	case "ask":
		response, err := newConversation(apiKey)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// printResponse(response)

		// Get the conversation_id from the response
		conversationID := response["conversation_id"]

		// get the question from the command line it will be everything after the
		// first argument
		question := os.Args[3:]
		if len(question) == 0 {
			fmt.Fprintln(os.Stderr, "question is required")
			os.Exit(1)
		}
		questionString := strings.Join(question, " ")
		url := "https://api.mendable.ai/v0/mendableChat"

		data := map[string]interface{}{
			"api_key":         apiKey,
			"question":        questionString,
			"history":         []interface{}{},
			"conversation_id": conversationID,
			"shouldStream":    false,
		}

		payload, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// fmt.Println(reflect.TypeOf(body))

		// store the body in a strings
		bodyString := string(body)

		// fmt.Println(bodyString)

		// get the "answer" from the body
		answer := strings.Split(bodyString, "answer")
		// we will get just the "text" from the answer and store it into an array
		text := strings.Split(answer[1], ":{\"text\":")

		// split the array at the "

		// create a string from the array
		textString := strings.Join(text, " ")

		// print type of text
		// fmt.Println(reflect.TypeOf(text))

		// // // print the answer
		// fmt.Println(textString)

		// remove the " quotation marks from the string
		textString = strings.ReplaceAll(textString, "\"", "")

		// remove the } from the string
		textString = strings.ReplaceAll(textString, "}", "")

		// drop/replace everything after the "soources" in the string (i have no clue wtf I am doin XD)
		textString = strings.Split(textString, ",sources")[0]

		// print the answer

		fmt.Println(textString)

	case "ingest":
		dataSource := os.Args[3]
		dataType := os.Args[4]

		if dataSource == "" || dataType == "" {
			fmt.Fprintln(os.Stderr, "url and Type are required")
			os.Exit(1)
		}
		response, err := ingestData(apiKey, dataSource, dataType)
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

func chat(apiKey string, question string, history []map[string]string, conversationID float64) (string, error) {
	url := "https://api.mendable.ai/v0/mendableChat"

	data := map[string]interface{}{
		"api_key":         apiKey,
		"question":        question,
		"history":         history,
		"conversation_id": conversationID,
		// "shouldStream":    false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// req.Header.Set("Accept", "text/event-stream")
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
