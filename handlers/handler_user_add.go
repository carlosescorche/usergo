package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/carlosescorche/usergo/db"
	"github.com/carlosescorche/usergo/types"
)

// HandlerUserAdd handles requests to register users
func HandlerUserAdd(w http.ResponseWriter, r *http.Request) {
	const op = "AddUser"

	var payload types.User
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		errInfo := ErrorInfo{Code: "errInvalidPayload", Op: op, Message: "The payload is invalid"}
		Error(w, errInfo, 400)
		return
	}

	// Create mongo client

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := db.Connect(ctx)

	// Validate payload

	errorPayload := url.Values{}

	if len(payload.Email) == 0 {
		errorPayload.Add("email", "The email is required")
	}

	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(payload.Email) {
		errorPayload.Add("email", "The email is invalid")
	}

	if _, check := db.UserCheckEmail(client, ctx, payload.Email, ""); check {
		errorPayload.Add("email", "The email is registered")
	}

	if len(payload.FirstName) == 0 {
		errorPayload.Add("firstname", "The first name is required")
	}

	if !regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(payload.FirstName) {
		errorPayload.Add("firstname", "The first name must contain alphabetic characters")
	}

	if len(payload.LastName) == 0 {
		errorPayload.Add("lastname", "The last name is required")
	}

	if !regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(payload.LastName) {
		errorPayload.Add("lastname", "The last name must contain alphabetic characters")
	}

	if len(payload.Password) < 6 {
		errorPayload.Add("password", "The password must be at least 6 characters")
	}

	if len(errorPayload) > 0 {
		errInfo := ErrorInfo{Code: "errInvalidPayload", Op: op, Message: errorPayload}
		Error(w, errInfo, 422)
		return
	}

	_, err = db.UserAdd(client, ctx, payload)
	if err != nil {
		errInfo := ErrorInfo{Code: "errInternal", Op: op, Message: "The user could not be inserted"}
		Error(w, errInfo, 400)
		return
	}

	Success(w, nil, http.StatusCreated)
}
