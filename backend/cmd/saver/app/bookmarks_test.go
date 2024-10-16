package app

import (
	"context"
	"testing"

	"github.com/natalie-richards/wine-app/graph/model"
	"github.com/natalie-richards/wine-app/pkg/util"
	"github.com/pashagolub/pgxmock/v4"
)

func Test_AddBookmarks(t *testing.T) {
	a := App{}
	ctx := context.Background()
	mockDB, err := pgxmock.NewPool()
	a.DBConn = mockDB
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()

	testData := model.AddBookmarkRequest{
		Name:     util.Ptr("Test Wine"),
		Grape:    util.Ptr("Cabernet Sauvignon"),
		Region:   util.Ptr("Willamette Valley"),
		Location: util.Ptr("Trader Joes"),
		Notes:    util.Ptr("test"),
		Image:    util.Ptr("google.com"),
		Username: util.Ptr("app_user"),
	}

	mockDB.ExpectExec("INSERT INTO app.bookmarks").
		WithArgs(testData.Name, testData.Grape, testData.Region, testData.Location, testData.Notes, testData.Username, testData.Image).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mockDB.ExpectCommit()

	if err = a.AddBookmark(ctx, &testData); err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}
}

// TODO: Add test for GetBookmarks
// TODO: Add tests for GraphQL endpoints. Mock resolver
