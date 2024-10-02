package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerSystemRoutes(r *mux.Router) {
	r.HandleFunc("/system/ping", a.handleHello)
	r.HandleFunc("/system/info", a.handleSystemInfo)
}

func (a *API) handleHello(w http.ResponseWriter, r *http.Request) {
	jsonStringResponse(w, http.StatusOK, "OK!")
}

func (a *API) handleSystemInfo(w http.ResponseWriter, r *http.Request) {
	serverMetadata := a.app.GetServerMetadata()

	bytes, err := json.Marshal(serverMetadata)
	if err != nil {
		errorResponse(w, err)
	}

	jsonStringResponse(w, http.StatusOK, string(bytes))
}
