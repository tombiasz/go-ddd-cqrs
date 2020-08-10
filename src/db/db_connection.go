// TODO: update pkg name
package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func NewDbConnection(dbUrl string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	return conn
}
