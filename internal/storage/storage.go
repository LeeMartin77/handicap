package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leemartin77/handicap/internal/config"
	"github.com/leemartin77/handicap/internal/storage/db"
)

type Storage interface {
	GetTestData(ctx context.Context) string
	GetQuerier() db.Querier
}

type postgresStorage struct {
	pool    *pgxpool.Pool
	querier db.Querier
}

// GetTestData implements [Storage].
func (p *postgresStorage) GetTestData(ctx context.Context) string {
	var res string
	err := p.pool.QueryRow(ctx, "select test_value from test_data").Scan(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// GetQuerier implements [Storage].
func (p *postgresStorage) GetQuerier() db.Querier {
	return p.querier
}

func NewStorage(ctx context.Context, cfg *config.Config) (Storage, error) {
	pl, err := pgxpool.New(ctx, cfg.PostgresUrl)
	if err != nil {
		return nil, err
	}

	err = runMigrations(ctx, cfg, pl)
	if err != nil {
		return nil, err
	}

	q := db.New(pl)

	return &postgresStorage{
		pool:    pl,
		querier: q,
	}, nil
}
