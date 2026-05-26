package database

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    // enable the mysql dialect
	_ "github.com/jinzhu/gorm/dialects/postgres" // enable the postgres dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // enable the sqlite3 dialect
)

const dbName = "../test/test.db"

// New creates a new wrapper for the gorm database framework.
func New(dialect, connection string, log bool) (*Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We normally don't need that much connections, so we limit them.

// Sqlite cannot handle concurrent operations well so limit to one connection.

// Turn on foreign keys.

func createDirectoryIfSqlite(dialect string, connection string) error {
	_ = "STUB: not implemented"
	return nil
}

// Database is a wrapper for the gorm framework.
type Database struct {
	DB *gorm.DB
}

// Close closes the gorm database connection.
func (d *Database) Close() error { _ = "STUB: not implemented"; return nil }
