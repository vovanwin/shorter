package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"

	"github.com/vovanwin/shorter/internal/app/model"
)

type Client interface {
	Begin(context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Db struct {
	pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *Db {
	return &Db{
		pool: pool,
	}
}

func (m *Db) GetLink(code string) (model.URLLink, error) {
	type Greeting struct {
		ID          string
		FirstName   string
		LastName    string
		DateOfBirth time.Time
	}
	var ps Greeting

	err := m.pool.QueryRow(context.Background(), "select * from public.test").Scan(
		&ps.ID,
		&ps.FirstName,
		&ps.LastName,
		&ps.DateOfBirth,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ps)
	return model.URLLink{}, nil
}

func (m *Db) AddLink(model model.URLLink) error {

	return nil
}

func (m *Db) GetLinksUser(user uuid.UUID) ([]model.UserURLLinks, error) {

	return []model.UserURLLinks{}, nil
}

func (m *Db) Ping() error {
	var greeting string
	err := m.pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil
}
