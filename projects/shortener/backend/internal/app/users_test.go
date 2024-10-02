package app

import (
	"backend/model"
	"context"
	"testing"
)

func TestGetProfile(t *testing.T) {
	th := SetupTestHelper(t)

	var (
		ctx = context.Background()

		emptyEmail = ""
		email      = "danielwetan.io@gmail.com"

		ctxWithEmptyEmail = context.WithValue(ctx, "email", emptyEmail)
		ctxWithEmail      = context.WithValue(ctx, "email", email)

		user = &model.User{
			ID:    1,
			Email: email,
		}
	)

	testCases := []struct {
		name     string
		request  context.Context
		response *model.User
		wantErr  bool
		mockFunc func()
	}{
		{
			name:     "#1. Get profile with invalid email",
			request:  ctxWithEmptyEmail,
			wantErr:  true,
			mockFunc: func() {},
		},
		{
			name:     "#2. Get profile success",
			request:  ctxWithEmail,
			response: user,
			wantErr:  false,
			mockFunc: func() {
				th.Store.EXPECT().GetUserByEmail(ctxWithEmail, email).Return(user, nil)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			response, err := th.App.GetProfile(test.request)
			if (err != nil) != test.wantErr {
				t.Errorf("expected error = %v, but got = %v", test.wantErr, err)
			}

			if response != nil && response != test.response {
				t.Errorf("expected = %v, but got = %v", test.response, response)
			}
		})
	}
}
