package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/migrate"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type pgConfig struct {
	Dsn string
}

// NewPgConfig creates new pg config instance
func NewPgConfig(dsn string) *pgConfig {
	return &pgConfig{
		Dsn: dsn,
	}
}

// NewClient
func NewClient(ctx context.Context, maxAttempts int, maxDelay time.Duration, cfg *pgConfig) (pool *pgxpool.Pool, err error) {
	dsn := cfg.Dsn

	err = DoWithAttempts(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pgxCfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Unable to parse config: %v\n", err)
		}

		pool, err = pgxpool.NewWithConfig(ctx, pgxCfg)
		if err != nil {
			log.Println("Failed to connect to postgres... Going to do the next attempt")

			return err
		}

		return nil
	}, maxAttempts, maxDelay)

	if err != nil {
		log.Fatal("All attempts are exceeded. Unable to connect to postgres")
	}

	//Заполнение БД

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Unable to acquire a database connection: %v\n", err)
	}
	migrateDatabase(ctx, conn.Conn())

	return pool, nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error

	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--

			continue
		}

		return nil
	}

	return err
}

func migrateDatabase(ctx context.Context, conn *pgx.Conn) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	migrator, err := migrate.NewMigrator(ctx, conn, "schema_versions")
	if err != nil {
		log.Fatalf("Unable to create a migrator: %v\n", err)
	}

	pathMigrations := os.DirFS("./migrations")

	err = migrator.LoadMigrations(pathMigrations)
	if err != nil {
		log.Fatalf("Unable to load migrations: %v\n", err)
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		log.Fatalf("Unable to migrate: %v\n", err)
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		log.Fatalf("Unable to get current schema version: %v\n", err)
	}

	log.Printf("Migration done. Current schema version: %v\n", ver)
}
