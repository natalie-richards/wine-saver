package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/natalie-richards/wine-app/graph/model"
)

func (a *App) GetBookmarks(ctx context.Context) ([]*model.Bookmark, error) {
	// TODO: Get User from context, validate user exists, and use username input to get bookmarks
	rows, err := a.DBConn.Query(ctx, "SELECT * from app.bookmarks WHERE username = 'app_user'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	collectRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Bookmark])
	if err != nil {
		log.Fatal(err)
	}
	var bookmarks []*model.Bookmark
	for _, bookmark := range collectRows {
		bookmarks = append(bookmarks, &bookmark)
	}

	return bookmarks, nil
}

func (a *App) AddBookmark(ctx context.Context, bookmark *model.AddBookmarkRequest) error {
	// TODO: Get User from context, validate user exists before saving bookmark
	query := `INSERT INTO app.bookmarks (name, grape, region, location, notes, username, image) VALUES (@name, @grape, @region, @location, @notes, @username, @image)`
	args := pgx.NamedArgs{
		"name":     bookmark.Name,
		"grape":    bookmark.Grape,
		"region":   bookmark.Region,
		"location": bookmark.Location,
		"notes":    bookmark.Notes,
		"username": bookmark.Username,
		"image":    bookmark.Image,
	}
	_, err := a.DBConn.Exec(ctx, query, args)
	if err != nil {
		// TODO: set up logging
		log.Fatal(err)
		return err
	}
	return nil
}
