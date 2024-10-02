package app

import (
	"backend/internal/store"
	"backend/model"
	"context"
)

type app struct {
	store store.Store
}

type App interface {
	GetServerMetadata() *model.ServerMetadata

	// Auth
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	Register(ctx context.Context, request *model.CreateUserRequest) error
	ClaimJWTToken(tokenString string) (*model.JWTClaims, error)

	// Users
	GetProfile(ctx context.Context) (*model.User, error)

	// Short Urls
	CreateShortUrl(ctx context.Context, request *model.CreateShortUrlRequest) (*model.CreateShortUrlResponse, error)
	GetShortUrl(ctx context.Context, request *model.GetShortUrlRequest) (*model.ShortUrl, error)
	GetShortUrls(ctx context.Context) ([]*model.ShortUrl, error)
}

func New(store store.Store) App {
	app := &app{
		store: store,
	}

	return app
}
