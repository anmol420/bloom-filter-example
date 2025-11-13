package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := successResponse{
		Status:  http.StatusOK,
		Message: "Server Ready!",
		Data:    nil,
	}
	if err := app.sendSuccessResponse(w, &data); err != nil {
		app.internalServerError(w, r, err)
	}
}
