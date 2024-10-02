package app

import (
	"backend/model"
	"backend/pkg"
	"context"
)

func (a *app) GetProfile(ctx context.Context) (*model.User, error) {
	email := ctx.Value("email").(string)
	if err := pkg.Validator.ValidateVar(email, "required"); err != nil {
		return nil, err
	}

	return a.store.GetUserByEmail(ctx, email)
}
