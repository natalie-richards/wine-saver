package gcp

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var pgxConn *pgx.Conn

func GetDBConn(ctx context.Context) *pgx.Conn {

	// TODO: Scrap .env and include credentials in Cloud Secret Manager for security
	err := godotenv.Load(".env")
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
		dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
		dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
		dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	)

	dsn := fmt.Sprintf("user=%s password=%s database=%s", dbUser, dbPwd, dbName)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	var opts []cloudsqlconn.Option
	d, err := cloudsqlconn.NewDialer(ctx, opts...)
	if err != nil {
		panic(err)
	}
	// Use the Cloud SQL connector to handle connecting to the instance.
	// This approach does *NOT* require the Cloud SQL proxy.
	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}
	dbURI := stdlib.RegisterConnConfig(config)
	sqlDb, err := sql.Open("pgx", dbURI)
	if err != nil {
		panic(err)
	}
	conn, err := sqlDb.Conn(ctx)
	if err != nil {
		// handle error from acquiring connection from DB pool
	}

	err = conn.Raw(func(driverConn any) error {
		pgxConn = driverConn.(*stdlib.Conn).Conn() // conn is a *pgx.Conn
		return nil
	})
	if err != nil {
		panic(err)
	}

	return pgxConn
}
