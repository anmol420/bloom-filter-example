package main

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (app *application) sendSuccessResponse(w http.ResponseWriter, res *successResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	return json.NewEncoder(w).Encode(res)
}

func (app *application) sendErrorResponse(w http.ResponseWriter, res *errorResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	return json.NewEncoder(w).Encode(res)
}
