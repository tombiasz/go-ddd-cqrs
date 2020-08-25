package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDbConnection(dbUrl string) *pgxpool.Pool {
	con, err := pgxpool.Connect(context.Background(), dbUrl)

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	return con
}
