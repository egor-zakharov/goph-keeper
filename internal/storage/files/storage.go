package files

import (
	"context"
	"database/sql"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
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

func (s *storage) Create(ctx context.Context, file models.FileData, userID string) (*models.FileData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	row := s.db.QueryRowContext(ctx,
		`INSERT INTO files(meta, file_name, user_id) VALUES ($1, $2, $3)  returning id`,
		file.Meta, file.Name, userID).Scan(&file.ID)

	var pgErr *pgconn.PgError
	if errors.As(row, &pgErr) {
		return nil, ErrConflict
	}

	return &file, nil
}

func (s *storage) Read(ctx context.Context, userID string) (*[]models.FileData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	var data []models.FileData

	rows, err := s.db.QueryContext(ctx,
		`SELECT id, meta, file_name FROM files WHERE user_id=$1 order by created_at`,
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

		var id, meta, fileName string

		err = rows.Scan(&id, &meta, &fileName)
		if err != nil {
			return nil, err
		}
		item := models.FileData{
			ID:   id,
			Meta: meta,
			Name: fileName,
		}
		data = append(data, item)
	}
	return &data, nil
}

func (s *storage) Delete(ctx context.Context, id string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`DELETE FROM files WHERE id=$1 and user_id=$2`,
		id, userID)
	if err != nil {
		return err
	}
	return nil
}
