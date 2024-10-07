package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Web3Payload struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

func (app *Config) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := app.Repo.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorJSON(w, errors.New("no record found"), nil, http.StatusBadRequest)
			return
		}

		app.errorJSON(w, err, nil, http.StatusInternalServerError)
		return
	}

	log.Println(users)

	payload := jsonResponse{
		Error:      false,
		StatusCode: http.StatusAccepted,
		Message:    "users retrieved successfully",
		Data:       users,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) MainnetMissedCheckpoint(w http.ResponseWriter, r *http.Request) {

	url := os.Getenv("MainnetMissedCheckpoint")
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	payload := jsonResponse{
		Error:      false,
		StatusCode: http.StatusAccepted,
		Message:    "results retrieved successfully",
		Data:       response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
func (app *Config) TestnetMissedCheckpoint(w http.ResponseWriter, r *http.Request) {

	url := os.Getenv("TestnetMissedCheckpoint")
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	payload := jsonResponse{
		Error:      false,
		StatusCode: http.StatusAccepted,
		Message:    "results retrieved successfully",
		Data:       response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) MainnetHeimdalBlockHeight(w http.ResponseWriter, r *http.Request) {

	url := os.Getenv("MainnetHeimdalBlockHeight")
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	payload := jsonResponse{
		Error:      false,
		StatusCode: http.StatusAccepted,
		Message:    "results retrieved successfully",
		Data:       response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
func (app *Config) TestnetHeimdalBlockHeight(w http.ResponseWriter, r *http.Request) {

	url := os.Getenv("TestnetHeimdalBlockHeight")
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	payload := jsonResponse{
		Error:      false,
		StatusCode: http.StatusAccepted,
		Message:    "results retrieved successfully",
		Data:       response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
func (app *Config) MainnetBorLatestBlockDetails(w http.ResponseWriter, r *http.Request) {

	body := Web3Payload{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{"latest", false},
		ID:      1,
	}

	// Convert the body struct to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Create the request using bytes.NewBuffer to pass the JSON payload as an io.Reader
	url := os.Getenv("MainnetBorLatestBlockDetails")
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Set headers
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	payload := jsonResponse{
		Error:      false,
		StatusCode: http.StatusAccepted,
		Message:    "results retrieved successfully",
		Data:       response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
func (app *Config) TestnetBorLatestBlockDetails(w http.ResponseWriter, r *http.Request) {

	body := Web3Payload{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{"latest", false},
		ID:      1,
	}

	// Convert the body struct to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Create the request using bytes.NewBuffer to pass the JSON payload as an io.Reader
	url := os.Getenv("TestnetBorLatestBlockDetails")
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Set headers
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	payload := jsonResponse{
		Error:      false,
		StatusCode: http.StatusAccepted,
		Message:    "results retrieved successfully",
		Data:       response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// func (app *Config) MissedCheckpoint(w http.ResponseWriter, r *http.Request) {

// 	users, err := app.Repo.GetAll()
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			app.errorJSON(w, errors.New("no record found"), nil, http.StatusBadRequest)
// 			return
// 		}

// 		app.errorJSON(w, err, nil, http.StatusInternalServerError)
// 		return
// 	}

// 	log.Println(users)

// 	payload := jsonResponse{
// 		Error:      false,
// 		StatusCode: http.StatusAccepted,
// 		Message:    "users retrieved successfully",
// 		Data:       users,
// 	}

// 	app.writeJSON(w, http.StatusAccepted, payload)
// }
