package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vovanwin/shorter/internal/app/config"
	"github.com/vovanwin/shorter/internal/app/model"
	"log"
	"os"
)

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

func (m *DB) GetLinkByLong(long string) (model.URLLink, error) {
	var ps model.URLLink
	var err error

	err = m.pool.QueryRow(context.Background(), "select id, code, long, user_id from public.url_link  where long = $1 LIMIT 1", long).Scan(
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

func (m *DB) AddLink(model model.URLLink) (string, error) {
	row := m.pool.QueryRow(context.Background(),
		"INSERT INTO public.url_link  (long, code, user_id) VALUES ($1, $2, $3) RETURNING id",
		model.Long, model.Code, model.UserID)

	var id uint64
	err := row.Scan(&id)

	if err != nil {
		s := err.Error()
		if s == "ERROR: duplicate key value violates unique constraint \"url_link_long_key\" (SQLSTATE 23505)" {

			return "23505", nil
		}

	}

	if err != nil {
		fmt.Printf("Unable to INSERT: %v\n", err)
		return "", err
	}
	return "", nil
}

func (m *DB) GetLinksUser(user uuid.UUID) ([]model.UserURLLinks, error) {
	var ps []model.UserURLLinks

	rows, err := m.pool.Query(context.Background(),
		"select id, code, long, user_id from public.url_link  where user_id = $1", user)

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
