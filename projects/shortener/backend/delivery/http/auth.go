package api

import (
	"backend/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", a.handleRegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", a.handleLogin).Methods(http.MethodPost)
}

func (a *API) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, err)
		return
	}

	var request model.CreateUserRequest
	err = json.Unmarshal(requestBody, &request)
	if err != nil {
		errorResponse(w, err)
		return
	}

	err = a.app.Register(r.Context(), &request)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonStringResponse(w, http.StatusOK, string(defaultSuccessResponse))
}

func (a *API) handleLogin(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, err)
		return
	}

	var loginData model.LoginRequest
	err = json.Unmarshal(requestBody, &loginData)
	if err != nil {
		errorResponse(w, err)
		return
	}

	loginResponse, err := a.app.Login(r.Context(), &loginData)
	if err != nil {
		errorResponse(w, err)
		return
	}

	response, err := json.Marshal(loginResponse)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonStringResponse(w, http.StatusOK, string(response))
}
