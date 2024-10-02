package gcp

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var dbpool *pgxpool.Pool

func GetDBPool(ctx context.Context) *pgxpool.Pool {
	if dbpool == nil {
		// TODO: Scrap .env and include credentials in Cloud Secret Manager for security
		err := godotenv.Load("../../.env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		mustGetenv := func(k string) string {
			v := os.Getenv(k)
			if v == "" {
				log.Fatalf("Fatal Error in connect_tcp.go: %s environment variable not set.", k)
			}
			return v
		}
		var (
			dbUser    = mustGetenv("DB_USER")
			dbPwd     = mustGetenv("DB_PASS")
			dbTCPHost = mustGetenv("INSTANCE_HOST")
			dbPort    = mustGetenv("DB_PORT")
			dbName    = mustGetenv("DB_NAME")
		)

		dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
			dbTCPHost, dbUser, dbPwd, dbPort, dbName)

		dbpool, err = pgxpool.New(ctx, dbURI)
		if err != nil {
			panic(err)
		}
	}
	return dbpool
}

func usePool(p *pgxpool.Pool) {
	dbpool = p
}
