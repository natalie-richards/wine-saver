package application

import "github.com/jackc/pgx/v5"

type Application struct {
	DBConn *pgx.Conn
}
