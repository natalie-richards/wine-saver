package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	DBConn *pgxpool.Pool
}

// Some serve logic
