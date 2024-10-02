package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerUserRoutes(r *mux.Router) {
	r.HandleFunc("/users/profile", a.secureRoutes(a.handleGetProfile)).Methods(http.MethodGet)
}

func (a *API) handleGetProfile(w http.ResponseWriter, r *http.Request) {
	getUserResponse, err := a.app.GetProfile(r.Context())
	if err != nil {
		errorResponse(w, err)
		return
	}

	response, err := json.Marshal(getUserResponse)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonStringResponse(w, http.StatusOK, string(response))
}
