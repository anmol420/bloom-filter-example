package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	app.sendErrorResponse(w, &errorResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	app.sendErrorResponse(w, &errorResponse{Status: http.StatusBadRequest, Message: "Bad Request Error"})
}

func (app *application) customError(w http.ResponseWriter, r *http.Request, status int, err error) {
	log.Printf("Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	app.sendErrorResponse(w, &errorResponse{Status: status, Message: err.Error()})
}