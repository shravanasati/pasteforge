package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getPostgresConnURI(user, password, hostname, port, dbname string) string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, password, hostname, port, dbname)
}

func NewConn(user, password, hostname, port, dbname string) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, getPostgresConnURI(user, password, hostname, port, dbname))
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewConnPool(user, password, hostname, port, dbname string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, getPostgresConnURI(user, password, hostname, port, dbname))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}
	return conn, nil
}
