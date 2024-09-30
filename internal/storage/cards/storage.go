package cards

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

func (s *storage) Create(ctx context.Context, card models.Card, userID string) (*models.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	row := s.db.QueryRowContext(ctx,
		`INSERT INTO cards(number, expiration_date, holder_name, cvv, user_id) VALUES ($1, $2, $3, $4, $5)  returning id`,
		card.Number, card.ExpirationDate, card.HolderName, card.CVV, userID).Scan(&card.ID)

	var pgErr *pgconn.PgError
	if errors.As(row, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return nil, ErrConflict
	}

	return &card, nil
}

func (s *storage) Read(ctx context.Context, userID string) (*[]models.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	var cards []models.Card

	rows, err := s.db.QueryContext(ctx,
		`SELECT id, number, expiration_date, holder_name, cvv FROM cards WHERE user_id=$1 order by created_at`,
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

		var id, number, expirationDate, holderName, cvv string

		err = rows.Scan(&id, &number, &expirationDate, &holderName, &cvv)
		if err != nil {
			return nil, err
		}
		card := models.Card{
			ID:             id,
			Number:         number,
			ExpirationDate: expirationDate,
			HolderName:     holderName,
			CVV:            cvv,
		}
		cards = append(cards, card)
	}
	return &cards, nil
}

func (s *storage) Update(ctx context.Context, card models.Card, userID string) (*models.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`UPDATE cards SET number=$1, expiration_date=$2, holder_name=$3, cvv=$4, updated_at=now() WHERE id=$5  and user_id=$6`,
		card.Number, card.ExpirationDate, card.HolderName, card.CVV, card.ID, userID)
	if err != nil {
		return nil, err
	}
	return &card, nil

}

func (s *storage) Delete(ctx context.Context, cardID string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	_, err := s.db.ExecContext(ctx,
		`DELETE FROM cards WHERE id=$1 and user_id=$2`,
		cardID, userID)
	if err != nil {
		return err
	}
	return nil
}
