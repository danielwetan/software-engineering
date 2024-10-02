package api

import (
	"backend/internal/app"
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	app app.App
}

func New(app app.App) *API {
	return &API{
		app: app,
	}
}

const (
	defaultSuccessResponse = "{}"
)

func jsonStringResponse(w http.ResponseWriter, code int, message string) {
	setResponseHeader(w, "Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func errorResponse(w http.ResponseWriter, err error) {
	setResponseHeader(w, "Content-Type", "application/json")

	// hardcoded for a while
	if err.Error() == "unauthorized" {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	errorResponse := model.ErrorResponse{Error: err.Error()}
	data, err := json.Marshal(errorResponse)
	if err != nil {
		data = []byte("{}")
	}

	_, _ = w.Write(data)
}

func setResponseHeader(w http.ResponseWriter, key string, value string) {
	header := w.Header()
	if header == nil {
		return
	}
	header.Set(key, value)
}
