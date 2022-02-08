package repo_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"github.com/stonelike/CleanGo/src/config"
)

func TestMain(m *testing.M) {

	dbInfo := config.CreateTestDBInfo()
	prepare(dbInfo)

	txdb.Register("txdb", "postgres", dbInfo.DSN)

	code := m.Run()
	os.Exit(code)
}

func prepare(dbInfo *config.DB) {
	db, err := sql.Open(dbInfo.Driver, dbInfo.DSN)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect(dbInfo.Driver),
		testfixtures.Directory("./fixtures"),
	)

	if err != nil {
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
