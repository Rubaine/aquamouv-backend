package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// SQLCtx is the context used for PostgreSQL operations.
var (
	SQLCtx  context.Context
	SQLConn *pgx.ConnConfig
)
