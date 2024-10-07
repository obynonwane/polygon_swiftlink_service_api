package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

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

	log.Println(response)

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

	log.Println(response)

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

	log.Println(response)

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

	log.Println(response)

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
