package app

import (
	"backend/internal/store/mockstore"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type TestHelper struct {
	App   App
	Store *mockstore.MockStore
}

func SetupTestHelper(t *testing.T) *TestHelper {
	ctrl := gomock.NewController(t)
	store := mockstore.NewMockStore(ctrl)

	mainApp := New(store)

	return &TestHelper{
		App:   mainApp,
		Store: store,
	}
}
