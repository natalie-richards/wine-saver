package graph

import (
	"github.com/natalie-richards/wine-app/cmd/saver/app"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	App *app.App
}
