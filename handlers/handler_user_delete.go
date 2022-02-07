package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/carlosescorche/usergo/db"
	"github.com/gorilla/mux"
)

// HandlerUserDelete handles requests to delete users
func HandlerUserDelete(w http.ResponseWriter, r *http.Request) {
	const op = "DeleteUser"

	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := db.Connect(ctx)

	_, check := db.UserCheckId(client, ctx, id)

	if !check {
		err := ErrorInfo{Code: "errNotFound", Message: "User not found", Op: op}
		Error(w, err, http.StatusNotFound)
		return

	}

	err := db.UserDelete(client, ctx, id)

	if err != nil {
		err := ErrorInfo{Code: "errInternal", Message: "There was a problem trying to delete the user", Op: op}
		Error(w, err, http.StatusInternalServerError)
		return
	}

	Success(w, "The user was deleted successfully", http.StatusOK)
}
