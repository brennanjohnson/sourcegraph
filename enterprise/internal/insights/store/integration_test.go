package store

import (
	"context"
	"database/sql"
	"os"
	"os/user"
	"testing"

	"github.com/sourcegraph/sourcegraph/internal/db/dbconn"
	"github.com/sourcegraph/sourcegraph/internal/db/dbutil"
	"github.com/sourcegraph/sourcegraph/internal/timeutil"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	getTimescaleDB := func(t testing.TB) *sql.DB {
		// Setup TimescaleDB for testing.
		username := ""
		if user, err := user.Current(); err == nil {
			username = user.Username
		}
		timescaleDSN := dbutil.PostgresDSN("codeinsights", username, os.Getenv)
		db, err := dbconn.New(timescaleDSN, "insights-test-"+t.Name())
		if err != nil {
			t.Fatalf("Failed to connect to codeinsights database: %s", err)
		}
		if err := dbconn.MigrateDB(db, "codeinsights"); err != nil {
			t.Fatalf("Failed to perform codeinsights database migration: %s", err)
		}
		return db
	}

	t.Run("Integration", func(t *testing.T) {
		ctx := context.Background()
		clock := timeutil.Now
		store := NewWithClock(getTimescaleDB(t), clock)
		t.Run("Insights", func(t *testing.T) { testInsights(t, ctx, store, clock) })
	})
}
