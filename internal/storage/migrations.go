package storage

import (
	"context"
	"embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leemartin77/handicap/internal/config"
)

//go:embed migrations/*.sql
var migrations embed.FS

func runMigrations(ctx context.Context, cfg *config.Config, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "CREATE TABLE IF NOT EXISTS migrations (version_number TEXT NOT NULL PRIMARY KEY, success BOOLEAN NOT NULL)")
	if err != nil {
		return err
	}
	mgs, err := migrations.ReadDir("migrations")
	if err != nil {
		return err
	}
	for _, mg := range mgs {
		if mg.IsDir() {
			return fmt.Errorf("critical error - directory in migrations embed")
		}
		mgkey := mg.Name()
		key := strings.Split(mgkey, "_")[0]
		if len(key) != 4 || strings.ContainsFunc(key, func(r rune) bool {
			return !unicode.IsNumber(r)
		}) {
			return fmt.Errorf("critical error - migration without 4 digit key (%s)", mgkey)
		}

		rw := pool.QueryRow(ctx, "SELECT * FROM migrations WHERE version_number = $1", key)
		rwer := rw.Scan()
		if rwer != nil && rwer != pgx.ErrNoRows {
			return rwer
		}
		if rwer != pgx.ErrNoRows {
			continue
		}
		fl, err := migrations.ReadFile("migrations/" + mg.Name())
		if err != nil {
			return err
		}
		str_mg := string(fl)
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return err
		}
		defer func() {
			tx.Rollback(ctx)
		}()
		_, err = tx.Exec(ctx, "INSERT INTO migrations (version_number, success) VALUES ($1, true)", key)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, str_mg)
		if err != nil {
			return err
		}
		err = tx.Commit(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
