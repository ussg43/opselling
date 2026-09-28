package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type SheetsUser struct {
	SheetsID int64 `db:"sheets_id"`
	Status   bool  `db:"status"`
}

type SheetsUserStore struct {
	db *sqlx.DB
}

func (s *SheetsUserStore) Create(ctx context.Context, u *SheetsUser) error {
	query := `
	INSERT INTO sheets (sheets_id, status)
	VALUES ($1, $2)`

	if _, err := s.db.ExecContext(
		ctx,
		query,
		&u.SheetsID,
		&u.Status,
	); err != nil {
		return err
	}

	return nil
}

func (s *SheetsUserStore) GetSheetID(ctx context.Context, tx *sqlx.Tx, idHash string) (*SheetsUser, error) {
	query := `
        SELECT sheets_id, status
        FROM sheets
        WHERE sheets_id = $1;
    `
	//Reminder: add id hashing
	user := &SheetsUser{}
	err := tx.QueryRowContext(
		ctx,
		query,
		idHash,
	).Scan(
		&user.SheetsID,
		&user.Status,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SheetsUserStore) Update(ctx context.Context, tx *sqlx.Tx, u *SheetsUser) error {
	query := `UPDATE sheets SET sheets_id = $1, status = $2`

	if _, err := tx.ExecContext(
		ctx,
		query,
		u.SheetsID,
		u.Status,
	); err != nil {
		return err
	}

	return nil
}

func (s *SheetsUserStore) DeleteUser(ctx context.Context, tx *sqlx.Tx, idHash string) error {
	return s.delete(ctx, tx, idHash)
}

func (s *SheetsUserStore) delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	query := `DELETE FROM sheets where sheets_id = $1`

	if _, err := tx.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}
