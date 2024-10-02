package sqlstore

import (
	"backend/model"
	"context"
	"database/sql"
	"errors"
)

const ShortUrlTableName = "short_urls"

func (s *SQLStore) CreateShortUrl(ctx context.Context, request *model.CreateShortUrlRequest) (*model.CreateShortUrlResponse, error) {
	query := `INSERT INTO ` + ShortUrlTableName + ` (user_id, target, shortcode) VALUES (?, ?, ?)`
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, model.ErrFailedToPrepareStatement
	}

	_, err = stmt.ExecContext(ctx,
		request.UserID,
		request.Target,
		request.ShortCode,
	)
	if err != nil {
		return nil, err
	}

	return &model.CreateShortUrlResponse{
		ShortCode: request.ShortCode,
	}, nil
}

func (s *SQLStore) GetShortUrl(ctx context.Context, request *model.GetShortUrlRequest) (*model.ShortUrl, error) {
	var shortUrl model.ShortUrl

	stmt, err := s.db.PrepareContext(ctx, `SELECT 
			shortcode,
			user_id,
			target,
			created_at
			FROM `+ShortUrlTableName+`
			WHERE shortcode = ?`)
	if err != nil {
		return nil, model.ErrFailedToPrepareStatement
	}

	err = stmt.QueryRowContext(ctx, request.ShortCode).Scan(
		&shortUrl.ShortCode,
		&shortUrl.UserID,
		&shortUrl.Target,
		&shortUrl.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, model.ErrDataNotFound
	}
	if err != nil {
		return nil, err
	}

	return &shortUrl, nil
}

// TODO: data filter
func (s *SQLStore) GetShortUrls(ctx context.Context) ([]*model.ShortUrl, error) {
	var results []*model.ShortUrl

	stmt, err := s.db.PrepareContext(ctx, `SELECT 
		shortcode,
		user_id,
		target,
		created_at
	FROM short_urls`)
	if err != nil {
		return nil, errors.New("failed to prepare statement")
	}

	rows, err := stmt.QueryContext(ctx)
	if err == sql.ErrNoRows {
		return nil, errors.New("data not found")
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var shortUrl model.ShortUrl
		err := rows.Scan(&shortUrl.ShortCode, &shortUrl.UserID, &shortUrl.Target, &shortUrl.CreatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, &shortUrl)
	}

	return results, nil
}
