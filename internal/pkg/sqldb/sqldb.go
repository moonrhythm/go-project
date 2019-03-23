package sqldb

import (
	"database/sql"

	"github.com/acoshift/pgsql"
	"github.com/acoshift/pgsql/pgctx"
	_ "github.com/lib/pq"

	"go-project/internal/pkg/config"
	"go-project/internal/pkg/middleware"
)

// Alias
var (
	Query      = pgctx.Query
	QueryRow   = pgctx.QueryRow
	Exec       = pgctx.Exec
	RunInTx    = pgctx.RunInTx
	NewContext = pgctx.NewContext
	Committed  = pgctx.Committed
	ErrAbortTx = pgsql.ErrAbortTx
)

const driver = "postgres"

var db *sql.DB

func configure(db *sql.DB) {
	db.SetMaxOpenConns(config.IntDefault("db_max_open_conns", 0))
	db.SetMaxIdleConns(config.IntDefault("db_max_idle_conns", 10))
	db.SetConnMaxLifetime(config.DurationDefault("db_conn_max_lifetime", 0))
}

// Open opens new database connection pool
func Open() error {
	var err error
	db, err = sql.Open(driver, config.String("db_url"))
	if err != nil {
		return err
	}
	configure(db)
	return nil
}

// Close closes database connection pool
func Close() error {
	return db.Close()
}

// Middleware returns a middleware that injects db into context
func Middleware() middleware.Middleware {
	return pgctx.Middleware(db)
}
