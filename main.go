package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type APIResponse struct {
	StatusCode int       `json:"status_code"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
	Path       string    `json:"path"`
	RequestID  string    `json:"request_id"`
}

type ResponseVariant struct {
	StatusCode int
	Status     string
	Messages   []string
}

var responses = []ResponseVariant{
	// Informational
	{100, "Continue", []string{
		"Request received, continuing processing.",
		"Headers received, awaiting body.",
		"Initial part of the request accepted.",
		"Please continue sending the request.",
	}},
	{101, "Switching Protocols", []string{
		"Protocol switch initiated.",
		"Server is switching protocols.",
		"Handshake complete for new protocol.",
		"Client requested a protocol change.",
	}},
	// Success
	{200, "OK", []string{
		"Request completed successfully.",
		"Everything is working as expected.",
		"Data fetched successfully.",
		"Operation executed correctly.",
	}},
	{201, "Created", []string{
		"New resource was created successfully.",
		"The user was successfully registered.",
		"Resource added to the database.",
		"Creation successful and acknowledged.",
	}},
	// Client errors
	{400, "Bad Request", []string{
		"Malformed request syntax.",
		"Invalid parameters sent.",
		"Check your input format.",
		"Request could not be understood.",
	}},
	{404, "Not Found", []string{
		"The resource does not exist.",
		"Endpoint not found.",
		"Nothing was found at this URL.",
		"No matching route.",
	}},
	// Server errors
	{500, "Internal Server Error", []string{
		"Something went wrong on our end.",
		"Unexpected condition encountered.",
		"We're fixing an internal issue.",
		"Oops! Server crashed temporarily.",
	}},
	{503, "Service Unavailable", []string{
		"Service is temporarily overloaded.",
		"Try again later, server busy.",
		"Service down for maintenance.",
		"System is currently unavailable.",
	}},
}

func generateRandomResponse(path string) APIResponse {
	rand.Seed(time.Now().UnixNano())
	resp := responses[rand.Intn(len(responses))]
	msg := resp.Messages[rand.Intn(len(resp.Messages))]

	return APIResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Message:    msg,
		Timestamp:  time.Now(),
		Path:       path,
		RequestID:  uuid.New().String(),
	}
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")
	count := 1
	if countStr != "" {
		if parsed, err := strconv.Atoi(countStr); err == nil && parsed > 0 && parsed <= 100 {
			count = parsed
		}
	}

	var result []APIResponse
	for i := 0; i < count; i++ {
		result = append(result, generateRandomResponse(r.URL.Path))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/api/random-response", responseHandler)

	port := ":9090"
	fmt.Println("Server running at http://localhost" + port)
	http.ListenAndServe(port, nil)
}
