package api

import (
	"backend/model"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) registerShortUrlRedirectionRoutes(r *mux.Router) {
	// simplify the redirection logic
	// for the real-world app we can separate this as a different service with custom domain
	r.HandleFunc("/s/{code}", a.handleGetShortUrl).Methods(http.MethodGet)
}

func (a *API) handleGetShortUrl(w http.ResponseWriter, r *http.Request) {
	request := &model.GetShortUrlRequest{
		ShortCode: mux.Vars(r)["code"],
	}

	getShortUrlResponse, err := a.app.GetShortUrl(r.Context(), request)
	if err != nil {
		errorResponse(w, err)
		return
	}

	if getShortUrlResponse.Target == "" {
		errorResponse(w, errors.New("invalid short code or URL not found"))
		return
	}

	http.Redirect(w, r, getShortUrlResponse.Target, http.StatusFound)
}
