package app

import (
	"backend/model"
	"backend/pkg"
	"context"
)

func (a *app) CreateShortUrl(ctx context.Context, request *model.CreateShortUrlRequest) (*model.CreateShortUrlResponse, error) {
	if err := pkg.Validator.ValidateStruct(request); err != nil {
		return nil, err
	}

	// set default to guest
	if request.UserID == 0 {
		request.UserID = 1
	}

	// dont check duplicate
	// keep the duplicate target

	request.ShortCode = pkg.GenerateShortCode()

	return a.store.CreateShortUrl(ctx, request)
}

func (a *app) GetShortUrl(ctx context.Context, request *model.GetShortUrlRequest) (*model.ShortUrl, error) {
	if err := pkg.Validator.ValidateStruct(request); err != nil {
		return nil, err
	}

	return a.store.GetShortUrl(ctx, request)
}

func (a *app) GetShortUrls(ctx context.Context) ([]*model.ShortUrl, error) {
	return a.store.GetShortUrls(ctx)
}
