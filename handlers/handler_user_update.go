package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/carlosescorche/usergo/db"
	"github.com/carlosescorche/usergo/types"
	"github.com/gorilla/mux"
)

// HandlerUserUpdate allows to handle requests to update users
func HandlerUserUpdate(w http.ResponseWriter, r *http.Request) {
	const op = "UpdateUser"

	vars := mux.Vars(r)
	id := vars["id"]

	var payload types.User
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		errInfo := ErrorInfo{Code: "errInvalidPayload", Op: op, Message: "The payload is invalid"}
		Error(w, errInfo, 400)
		return
	}

	// Create the mongo client

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := db.Connect(ctx)

	// Verify if the user is registered

	_, check := db.UserCheckId(client, ctx, id)

	if !check {
		errInfo := ErrorInfo{Code: "errNotFound", Message: "User not found", Op: op}
		Error(w, errInfo, http.StatusNotFound)
		return
	}

	// Validate payload

	errorPayload := url.Values{}

	if len(payload.Email) > 0 {
		if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(payload.Email) {
			errorPayload.Add("email", "The email is invalid")
		}

		if _, check := db.UserCheckEmail(client, ctx, payload.Email, id); check {
			errorPayload.Add("email", "The email is registered")
		}
	}

	if len(payload.FirstName) > 0 {
		if !regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(payload.FirstName) {
			errorPayload.Add("firstname", "The first name must contain alphabetic characters")
		}
	}

	if len(payload.LastName) > 0 {
		if !regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(payload.LastName) {
			errorPayload.Add("lastname", "The last name must contain alphabetic characters")
		}

	}

	if len(payload.Password) > 0 && len(payload.Password) < 6 {
		errorPayload.Add("password", "The password must be at least 6 characters")
	}

	if len(errorPayload) > 0 {
		errInfo := ErrorInfo{Code: "errInvalidPayload", Op: op, Message: errorPayload}
		Error(w, errInfo, 422)
		return
	}

	if len(payload.Password) > 0 {
		payload.Password, _ = db.PasswordEncrypt(payload.Password)
	}

	err = db.UserUpdate(client, ctx, id, payload)

	if err != nil {
		log.Println(err.Error())

		err := ErrorInfo{Code: "errInternal", Message: "There was a problem trying to update the user", Op: op}
		Error(w, err, http.StatusInternalServerError)
		return
	}

	Success(w, "The user was updated successfully", http.StatusOK)
}
