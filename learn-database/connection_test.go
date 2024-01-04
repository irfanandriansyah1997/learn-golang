package learndatabase

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestDBConnection(t *testing.T) {
	db, err := sql.Open("mysql", "my_user:my_password@tcp(localhost:3306)/my_database")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	defer db.Close()

}
