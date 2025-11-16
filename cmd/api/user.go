package main

import (
	"errors"
	"net/http"

	"github.com/anmol420/bloom-filter-example/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreateUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=20"`
}

type createUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
	// Database calls
	userFound, err := app.store.Users.FindUser(r.Context(), payload.Username, payload.Email)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if userFound {
		app.customError(w, r, http.StatusBadRequest, errors.New("User Already Exists"))
		return
	}
	createdUser := &store.User{Username: payload.Username, Email: payload.Email, Password: payload.Password}
	if err := app.store.Users.Create(r.Context(), createdUser); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	// TODO: Cache update -> Add username to cache
	if err := app.sendSuccessResponse(w, &successResponse{Status: http.StatusCreated, Data: &createUserResponse{Username: payload.Username, Email: payload.Email, Password: payload.Password}, Message: "User created"}); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) searchUserHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := app.store.Users.SearchByUsername(r.Context(), username)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if user == nil {
		app.customError(w, r, http.StatusNotFound, errors.New("User Not Found"))
		return
	}
	data := map[string]any {
		"username": user.Username,
		"email": user.Email,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	}
	if err := app.sendSuccessResponse(w, &successResponse{Status: http.StatusOK, Data: data, Message: "User Searched"}); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
