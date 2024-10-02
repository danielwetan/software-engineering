package app

import (
	"backend/model"
	"runtime"
)

func (a *app) GetServerMetadata() *model.ServerMetadata {
	var dbType string
	var dbVersion string
	if a != nil && a.store != nil {
		dbType = a.store.DBType()
		dbVersion = a.store.DBVersion()
	}

	return &model.ServerMetadata{
		Version:   model.CurrentVersion,
		DBType:    dbType,
		DBVersion: dbVersion,
		OSType:    runtime.GOOS,
		OSArch:    runtime.GOARCH,
	}
}
