package api

import (
	"backend/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerShortUrlRoutes(r *mux.Router) {
	// r.HandleFunc("/s", a.secureRoutes(a.handleCreateShortUrl)).Methods(http.MethodPost)
	r.HandleFunc("/s/public", a.handleCreateShortUrlPublic).Methods(http.MethodPost)
	r.HandleFunc("/s/public", a.handleGetShortUrls).Methods(http.MethodGet)
}

func (a *API) handleCreateShortUrlPublic(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, err)
		return
	}

	var createShortUrlRequest model.CreateShortUrlRequest
	err = json.Unmarshal(requestBody, &createShortUrlRequest)
	if err != nil {
		errorResponse(w, err)
		return
	}

	createShortUrlResponse, err := a.app.CreateShortUrl(r.Context(), &createShortUrlRequest)
	if err != nil {
		errorResponse(w, err)
		return
	}

	response, err := json.Marshal(createShortUrlResponse)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonStringResponse(w, http.StatusOK, string(response))
}

func (a *API) handleGetShortUrls(w http.ResponseWriter, r *http.Request) {
	getShortUrlsResponse, err := a.app.GetShortUrls(r.Context())
	if err != nil {
		errorResponse(w, err)
		return
	}

	response, err := json.Marshal(getShortUrlsResponse)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonStringResponse(w, http.StatusOK, string(response))
}
