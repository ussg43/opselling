package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	Players interface {
		Create(context.Context, *Player) error
		Update(context.Context, *sqlx.Tx, *Player) error
		Delete(context.Context, *sqlx.Tx, int64) error
		GetPlayerByID(context.Context, *sqlx.Tx, int64) (*Player, error)
	}

	SheetsUsers interface {
		Create(context.Context, *SheetsUser) error
		GetSheetID(context.Context, *sqlx.Tx, string) (*SheetsUser, error)
		DeleteUser(context.Context, *sqlx.Tx, string) error
	}
}

func NewStorage(db *sqlx.DB) Storage {
	return Storage{
		Players:     &PlayerStore{db},
		SheetsUsers: &SheetsUserStore{db},
	}
}

func withTx(ctx context.Context, db *sqlx.DB, fn func(tx *sqlx.Tx) error) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
