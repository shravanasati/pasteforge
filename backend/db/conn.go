package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getPostgresConnURI(user, password, dbname string) string {
	return fmt.Sprintf("postgres://%v:%v@localhost:5432/%v", user, password, dbname)
}

func NewConn(user, password, dbname string) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, getPostgresConnURI(user, password, dbname))
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewConnPool(user, password, dbname string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, getPostgresConnURI(user, password, dbname))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}
	return conn, nil
}
