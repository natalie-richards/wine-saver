package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
	cloudstorage "github.com/natalie-richards/wine-app/pkg/cloud-storage"
)

type Application struct {
	DBConn      *pgxpool.Pool
	CloudClient cloudstorage.Client
}

// Some serve logic
