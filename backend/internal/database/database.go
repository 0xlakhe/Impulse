package database

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Database {
	return &Database{Pool: pool}
}

func IsNotFound(err error) bool{
	return errors.Is(err,pgx.ErrNoRows)
}