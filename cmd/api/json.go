package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	max_bytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(max_bytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}
