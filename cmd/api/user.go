package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CreateUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
}

type createUserResponse struct {
	Username string `json:"username"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload
	if err := app.readJson(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}
	// TODO: Cache Check
	// TODO: Database calls -> 1. Create User
	// TODO: Cache update -> Add username to cache
	if err := app.sendSuccessResponse(w, &successResponse{Status: http.StatusCreated, Data: &createUserResponse{Username: payload.Username}, Message: "User created"}); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) searchUserHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	data := map[string]any{
		"username": username,
		"found": true,
	}
	if err := app.sendSuccessResponse(w, &successResponse{Status: http.StatusOK, Data: data, Message: "User Searched"}); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}