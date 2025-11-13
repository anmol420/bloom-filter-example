package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	app.sendErrorResponse(w, &errorResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
}