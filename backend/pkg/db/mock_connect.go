package gcp

import (
	"github.com/pashagolub/pgxmock/v4"
)

// TODO: Move pgx mocks in to their own package

type PgxConnMock struct {
	pgxmock.PgxConnIface
}
