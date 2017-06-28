// Original: https://github.com/mattn/go-sqlite3/blob/ed211752885a3e9a33d39e12e8d0fa46c12d85fa/_example/simple/simple.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// https://golang.org/pkg/database/sql/#Open
	// Setup a connection to a database.  The first argument is the
	// driver and the second argument is the source syntax expected
	// by the driver.  The database and error are returned.
	sqlite3_file := "test.db"
	db, open_err := sql.Open("sqlite3", sqlite3_file)
	if open_err != nil {
		log.Fatal(open_err)
	}
	// Closes the database when the function ends.  This will
	// frequently be done at a scope that will stay around longer,
	// like here in main().
	defer db.Close()

	// Here we create the table using `Exec()`, which does not
	// return rows like `Query()` does.  In fact, I don't even know
	// what the result is except that it is a struct with only
	// `{id, changes int64}`.  These two values trace back to a couple
	// lines of embedded C.
	// ```
	//    *rowid = (long long) sqlite3_last_insert_rowid(db);
	//    *changes = (long long) sqlite3_changes(db);
	// ```
	// For more see: https://sqlite.org/c3ref/last_insert_rowid.html
	// and https://sqlite.org/c3ref/changes.html.
	//
	// I'll just ignore the first return for now.
	table_def := "CREATE TABLE reminders (id INTEGER PRIMARY KEY, created_at DATETIME, remind_at DATETIME, message TEXT, reminded BOOLEAN)"
	_, create_err := db.Exec(table_def)
	if create_err != nil {
		log.Fatal(create_err)
	}

	fmt.Println("database created: " + sqlite3_file)
}
