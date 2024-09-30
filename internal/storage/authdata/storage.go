package authdata

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

func (s *storage) Create(ctx context.Context, authData models.AuthData, userID string) (*models.AuthData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	row := s.db.QueryRowContext(ctx,
		`INSERT INTO auth(meta, login, password, user_id) VALUES ($1, $2, $3, $4)  returning id`,
		authData.Meta, authData.Login, authData.Password, userID).Scan(&authData.ID)

	var pgErr *pgconn.PgError
	if errors.As(row, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return nil, ErrConflict
	}

	return &authData, nil
}

func (s *storage) Read(ctx context.Context, userID string) (*[]models.AuthData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	var data []models.AuthData

	rows, err := s.db.QueryContext(ctx,
		`SELECT id, meta, login, password FROM auth WHERE user_id=$1 order by created_at`,
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

		var id, meta, login, password string

		err = rows.Scan(&id, &meta, &login, &password)
		if err != nil {
			return nil, err
		}
		card := models.AuthData{
			ID:       id,
			Meta:     meta,
			Login:    login,
			Password: password,
		}
		data = append(data, card)
	}
	return &data, nil
}

func (s *storage) Update(ctx context.Context, authData models.AuthData, userID string) (*models.AuthData, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`UPDATE auth SET meta=$1, login=$2, password=$3, updated_at=now() WHERE id=$4  and user_id=$5`,
		authData.Meta, authData.Login, authData.Password, authData.ID, userID)
	if err != nil {
		return nil, err
	}
	return &authData, nil

}

func (s *storage) Delete(ctx context.Context, id string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`DELETE FROM auth WHERE id=$1 and user_id=$2`,
		id, userID)
	if err != nil {
		return err
	}
	return nil
}
