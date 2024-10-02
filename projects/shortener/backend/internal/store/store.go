package store

import (
	"backend/model"
	"context"
)

type Store interface {
	// General
	DBType() string
	DBVersion() string

	// Users
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, request *model.CreateUserRequest) error

	// Short Url
	CreateShortUrl(ctx context.Context, request *model.CreateShortUrlRequest) (*model.CreateShortUrlResponse, error)
	GetShortUrl(ctx context.Context, request *model.GetShortUrlRequest) (*model.ShortUrl, error)
	GetShortUrls(ctx context.Context) ([]*model.ShortUrl, error)
}
