package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO: Connection interface?
type DbConnectionFactory interface {
	GetConnection() *pgxpool.Pool
}

type dbConnectionFactory struct {
	connectionString string
}

func NewDbConnectionFactory(connectionString string) *dbConnectionFactory {
	return &dbConnectionFactory{connectionString}
}

func (d dbConnectionFactory) GetConnection() *pgxpool.Pool {
	con, err := pgxpool.Connect(context.Background(), d.connectionString)

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	return con
}
