package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Player struct {
	ID int64 `json:"player_id"`
	Name string `json:"name"`
	Price uint64 `json:"price"`
}

type PlayerStore struct {
	db *sqlx.DB
}

func (s *PlayerStore) Create(ctx context.Context, p *Player) error {
	query := `
	INSERT INTO sheets (id, name, price)
	VALUES ($1, $2, $3)`

	if _, err := s.db.ExecContext(
		ctx,
		query,
		&p.ID,
		&p.Name,
		&p.Price,
	); err != nil {
		return err
	}

	return nil
}

func (s *PlayerStore) GetPlayerByID(ctx context.Context, tx *sqlx.Tx, id int64) (*Player, error) {
	query := `
        SELECT id, name, price
        FROM players
        WHERE sheets_id = $1;
    `
	
	player := &Player{}
	err := tx.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&player.ID,
		&player.Name,
		&player.Price,
	)

	if err != nil {
		return nil, err
	}

	return player, nil
}

func (s *PlayerStore) Update(ctx context.Context, tx *sqlx.Tx, p *Player) error {
	query := `UPDATE players SET price = $1`

	if _, err := tx.ExecContext(
		ctx,
		query,
		p.Price,
	); err != nil {
		return err
	}

	return nil
}

func (s *PlayerStore) Delete(ctx context.Context, tx *sqlx.Tx, id int64) error {
	return s.delete(ctx, tx, id)
}

func (s *PlayerStore) delete(ctx context.Context, tx *sqlx.Tx, id int64) error {
	query := `DELETE FROM player where id = $1`

	if _, err := tx.ExecContext(ctx, query, id); err != nil {
		return err
	}

	return nil
}
