package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/natalie-richards/wine-app/graph/model"
)

func (a *App) GetBookmarks(ctx context.Context) ([]*model.Bookmark, error) {
	// TODO: Get User from context, validate user exists, and use it to get bookmarks
	rows, err := a.DBConn.Query(ctx, "select * from app.bookmarks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bookmarks, err := pgx.CollectRows(rows, pgx.RowToStructByName[*model.Bookmark])
	if err != nil {
		// TODO: set up logging
		log.Fatal(err)
		return nil, err
	}

	return bookmarks, nil
}

func (a *App) AddBookmark(ctx context.Context, bookmark *model.AddBookmarkRequest) error {
	// TODO: Get User from context, validate user exists before saving bookmark
	query := `INSERT INTO bookmarks (name, region, location) VALUES (@name, @region, @location)`
	args := pgx.NamedArgs{
		"name":     bookmark.Name,
		"region":   bookmark.Region,
		"location": bookmark.Location,
	}
	_, err := a.DBConn.Exec(ctx, query, args)
	if err != nil {
		// TODO: set up logging
		log.Fatal(err)
		return err
	}
	return nil
}
