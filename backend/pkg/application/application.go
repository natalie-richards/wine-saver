package application

import gcp "github.com/natalie-richards/wine-app/pkg/db"

type Application struct {
	DBConn gcp.PgxConnIface
}
