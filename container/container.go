package container

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	dbReaderOnce sync.Once
	dbWriterOnce sync.Once
	dbReader     *sqlx.DB
	dbWriter     *sqlx.DB
)

// SetWriter sets a db instance for the read/write operation
func SetWriter(db *sqlx.DB) {
	dbWriterOnce.Do(func() {
		dbWriter = db
	})
}

// SetReader sets a db instance for the read/write operation
func SetReader(db *sqlx.DB) {
	dbReaderOnce.Do(func() {
		dbReader = db
	})
}

// GetWriter gets a db instance for the read/write operation
func GetWriter() *sqlx.DB {
	return dbWriter
}

// GetReader gets a db instance for the read/write operation
func GetReader() *sqlx.DB {
	return dbReader
}
