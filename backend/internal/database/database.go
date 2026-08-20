package database

import "github.com/jackc/pgx/v5/pgxpool"


type Database struct{	
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Database{
	return &Database{Pool: pool}
}