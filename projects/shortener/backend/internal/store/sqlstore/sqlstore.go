package sqlstore

import (
	"backend/internal/store"
	"backend/model"
	"database/sql"
)

type SQLStore struct {
	db *sql.DB
}

func New(db *sql.DB) (store.Store, error) {
	return &SQLStore{
		db: db,
	}, nil
}

func (s *SQLStore) DBType() string {
	return model.DBTypeMySQL
}

func (s *SQLStore) DBVersion() string {
	var version string
	err := s.db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return ""
	}

	return version
}
