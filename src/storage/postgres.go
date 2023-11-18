package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgresql struct {
	db *pgxpool.Pool
}

func (s *Postgresql) Open(c string) (err error) {
	pool, err := pgxpool.New(context.Background(), c)
	if err != nil {
		return
	}
	s.db = pool
	return
}

func (s *Postgresql) Close() {
	s.db.Close()
}

func (s *Postgresql) Exec(qs string) (err error) {
	ctx := context.Background()
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, qs)
	if err != nil {
		return
	}
	return
}

func (s *Postgresql) Begin(qs string, d []string) (err error) {
	ctx := context.Background()
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return
	}
	defer conn.Release()

	var txoptions pgx.TxOptions
	tx, err := conn.BeginTx(ctx, txoptions)
	if err != nil {
		return
	}
	defer func() {
		err = tx.Rollback(ctx)
	}()
	return
}

func (s *Postgresql) Query(qs string) (result []string, err error) {
	return
}

func (s *Postgresql) Prepare(qs string) (result string, err error) {
	return
}
