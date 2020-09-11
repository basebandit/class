// Package tests contains supporting code for running tests.
package tests

import (
	"testing"
	"time"

	"github.com/ardanlabs/service/business/data/schema"
	"github.com/ardanlabs/service/foundation/database"
	"github.com/jmoiron/sqlx"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

// Configuration for running tests.
const (
	dbImage = "postgres:11.1-alpine"
	dbPort  = "5432"
	AdminID = "5cf37266-3473-4006-984f-9325122678b7"
	UserID  = "45b5fbd3-755f-4379-8f07-a58d4a30fa2f"
)

// NewUnit creates a test database inside a Docker container. It creates the
// required table structure but the database is otherwise empty. It returns
// the database to use as well as a function to call at the end of the test.
func NewUnit(t *testing.T) (*sqlx.DB, func()) {
	c := startContainer(t, dbImage, dbPort)

	cfg := database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       c.Host,
		Name:       "postgres",
		DisableTLS: true,
	}
	db, err := database.Open(cfg)
	if err != nil {
		t.Fatalf("opening database connection: %v", err)
	}

	t.Log("waiting for database to be ready ...")

	// Wait for the database to be ready. Wait 100ms longer between each attempt.
	// Do not try more than 20 times.
	var pingError error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
	}

	if pingError != nil {
		dumpContainerLogs(t, c.ID)
		stopContainer(t, c.ID)
		t.Fatalf("database never ready: %v", pingError)
	}

	if err := schema.Migrate(db); err != nil {
		stopContainer(t, c.ID)
		t.Fatalf("migrating error: %s", err)
	}

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		db.Close()
		stopContainer(t, c.ID)
	}

	return db, teardown
}
