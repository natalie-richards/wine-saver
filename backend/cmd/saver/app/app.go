package app

import (
	"context"

	application "github.com/natalie-richards/wine-app/pkg/application"
	gcp "github.com/natalie-richards/wine-app/pkg/db"
)

type App struct {
	application.Application
}

func (a *App) Init() {
	a.DBConn = gcp.GetDBPool(context.Background())
}
