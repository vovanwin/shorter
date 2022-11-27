package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/model"
	"log"
	"os"
)

type Client interface {
	Begin(context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type DB struct {
	pool   *pgxpool.Pool
	Config config.Config
}

func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{
		pool: pool,
	}
}

func (m *DB) GetLink(code string) (model.URLLink, error) {
	var ps model.URLLink
	var err error

	err = m.pool.QueryRow(context.Background(), "select id, code, long, user_id from public.url_link  where code = $1 LIMIT 1", code).Scan(
		&ps.ID,
		&ps.Code,
		&ps.Long,
		&ps.UserID,
	)
	fmt.Print(ps)
	if err != nil {
		err = errors.New("ссылка не найдена")
	}

	return ps, err
}

func (m *DB) AddLink(model model.URLLink) error {
	row := m.pool.QueryRow(context.Background(),
		"INSERT INTO public.url_link  (long, code, user_id) VALUES ($1, $2, $3) RETURNING id",
		model.Long, model.Code, model.UserID)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		fmt.Printf("Unable to INSERT: %v\n", err)
		return err
	}
	return nil
}

func (m *DB) GetLinksUser(user uuid.UUID) ([]model.UserURLLinks, error) {
	var ps []model.UserURLLinks

	rows, err := m.pool.Query(context.Background(), "select id, code, long, user_id from public.url_link  where user_id = $1", user)

	if err != nil {
		return []model.UserURLLinks{}, err

	}
	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		var url model.UserURLLinks
		// convert DB types to Go types
		url.ShortLink = fmt.Sprintf("%s/%s", m.Config.GetConfig().ServerAddress, values[1].(string))
		url.Long = values[2].(string)
		ps = append(ps, url)

	}

	return ps, nil
}

func (m *DB) Ping() error {
	err := m.pool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil
}
