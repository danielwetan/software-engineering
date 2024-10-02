package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	public := r.PathPrefix("/").Subrouter()
	a.registerShortUrlRedirectionRoutes(public)

	v1 := r.PathPrefix("/v1").Subrouter()
	a.registerSystemRoutes(v1)
	a.registerAuthRoutes(v1)
	a.registerUserRoutes(v1)
	a.registerShortUrlRoutes(v1)
}

func (a *API) secureRoutes(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			errorResponse(w, errors.New("unauthorized"))
			return
		}

		jwtClaims, err := a.app.ClaimJWTToken(authorization)
		if err != nil {
			errorResponse(w, errors.New("unauthorized"))
			return
		}

		// keep plain string for debugging purpose
		ctx := context.WithValue(r.Context(), "user_id", jwtClaims.UserID)
		ctx = context.WithValue(ctx, "email", jwtClaims.Email)
		handler(w, r.WithContext(ctx))
	}
}
