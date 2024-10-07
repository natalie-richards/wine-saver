package app

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/natalie-richards/wine-app/graph/model"
)

func (a *App) GetBookmarks(ctx context.Context) ([]*model.Bookmark, error) {
	// TODO: Get User from context, validate user exists, and use id to get bookmarks
	rows, err := a.DBConn.Query(ctx, "select * from app.bookmarks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var bookmarks []*model.Bookmark
	for rows.Next() {
		bookmark := model.Bookmark{}
		err := rows.Scan(&bookmark.Name, &bookmark.Region, &bookmark.Grape, &bookmark.Location, &bookmark.Notes, &bookmark.Username, &bookmark.Image)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		bookmarks = append(bookmarks, &bookmark)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
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
