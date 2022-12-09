package db

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	host := os.Getenv("DIGIKAM_DB_HOST")
	port := os.Getenv("DIGIKAM_DB_PORT")
	username := os.Getenv("DIGIKAM_DB_USER")
	password := os.Getenv("DIGIKAM_DB_USER_PASSWORD")
	dbName := os.Getenv("DIGIKAM_DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	var err error

	// create a database object which can be used
	// to connect with database.
	dbCon, err := sql.Open("mysql", connectionString)

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	testDb(dbCon)

	return dbCon
}

func testDb(dbCon *sql.DB) {
	err := dbCon.Ping()
	if err != nil {
		panic(err)
	}

	// test sql query
	_, err = dbCon.Query("SELECT id, name FROM digikam.Images limit 1")

	if err != nil {
		panic(err)
	}
}
