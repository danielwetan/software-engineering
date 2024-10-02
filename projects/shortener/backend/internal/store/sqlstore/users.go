package sqlstore

import (
	"backend/model"
	"context"
	"database/sql"
)

const UserTableName = "users"

func (s *SQLStore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	stmt, err := s.db.PrepareContext(ctx, `SELECT 
			id, 
			email,
			password,
			created_at
			FROM `+UserTableName+`
			WHERE email = ?`)
	if err != nil {
		return nil, model.ErrFailedToPrepareStatement
	}

	err = stmt.QueryRowContext(ctx, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, model.ErrDataNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *SQLStore) CreateUser(ctx context.Context, request *model.CreateUserRequest) error {
	query := `INSERT INTO ` + UserTableName + ` (email, password) VALUES (?, ?)`
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return model.ErrFailedToPrepareStatement
	}

	_, err = stmt.ExecContext(ctx,
		request.Email,
		request.Password,
	)

	return err
}
