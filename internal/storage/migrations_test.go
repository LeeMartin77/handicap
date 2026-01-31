package storage_test

import (
	"context"
	"log"
	"testing"

	"github.com/leemartin77/handicap/internal/config"
	"github.com/leemartin77/handicap/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func Test_StorageInitialises(t *testing.T) {
	ctx := context.Background()

	dbName := "postgres_test"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"postgres:13.23-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
	}
	require.NoError(t, err)

	strg, err := storage.NewStorage(ctx, &config.Config{
		PostgresUrl: postgresContainer.MustConnectionString(ctx),
	})

	require.NoError(t, err)

	td := strg.GetTestData(ctx)

	assert.Equal(t, "hello there", td)
}
