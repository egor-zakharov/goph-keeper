package textdata

import (
	"context"
	"database/sql"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"time"
)

const timeOut = 500 * time.Millisecond

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) Create(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	row := s.db.QueryRowContext(ctx,
		`INSERT INTO conf(meta, text, user_id) VALUES ($1, $2, $3)  returning id`,
		textData.Meta, textData.Text, userID).Scan(&textData.ID)

	var pgErr *pgconn.PgError
	if errors.As(row, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return nil, ErrConflict
	}

	return &textData, nil
}

func (s *storage) Read(ctx context.Context, userID string) (*[]models.TextData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	var data []models.TextData

	rows, err := s.db.QueryContext(ctx,
		`SELECT id, meta, text FROM conf WHERE user_id=$1 order by created_at`,
		userID)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var id, meta, text string

		err = rows.Scan(&id, &meta, &text)
		if err != nil {
			return nil, err
		}
		item := models.TextData{
			ID:   id,
			Meta: meta,
			Text: text,
		}
		data = append(data, item)
	}
	return &data, nil
}

func (s *storage) Update(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`UPDATE conf SET meta=$1, text=$2, updated_at=now() WHERE id=$3  and user_id=$4`,
		textData.Meta, textData.Text, textData.ID, userID)
	if err != nil {
		return nil, err
	}
	return &textData, nil

}

func (s *storage) Delete(ctx context.Context, id string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`DELETE FROM conf WHERE id=$1 and user_id=$2`,
		id, userID)
	if err != nil {
		return err
	}
	return nil
}
