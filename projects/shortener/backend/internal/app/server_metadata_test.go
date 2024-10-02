package app

import (
	"backend/model"
	"reflect"
	"runtime"
	"testing"
)

func TestGetServerMetadata(t *testing.T) {
	th := SetupTestHelper(t)

	const (
		dbType    = "mysql"
		dbVersion = "8.0.32"
	)

	want := &model.ServerMetadata{
		Version:   model.CurrentVersion,
		DBType:    dbType,
		DBVersion: dbVersion,
		OSType:    runtime.GOOS,
		OSArch:    runtime.GOARCH,
	}

	t.Run("#1. Success get server metadata", func(t *testing.T) {
		th.Store.EXPECT().DBType().Return(dbType)
		th.Store.EXPECT().DBVersion().Return(dbVersion)

		metadata := th.App.GetServerMetadata()
		if !reflect.DeepEqual(metadata, want) {
			t.Errorf("expected %v, but got %v", want, metadata)
		}
	})
}
