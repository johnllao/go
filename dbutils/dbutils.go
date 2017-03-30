package dbutils

import (
	"database/sql"
	"log"
)

// Database encapsulates database operation
type Database struct {
	driver     string
	connection string
}

// CreateDb create instance of the DB
func CreateDb(driver string, connection string) Database {
	db := Database{driver, connection}
	return db
}

// Query gets the query results from the database
func (d *Database) Query(qry string, factory func(*sql.Rows) interface{}, args ...interface{}) <-chan interface{} {

	ch := make(chan interface{})

	go func() {
		db, err := sql.Open(d.driver, d.connection)
		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}
		defer db.Close()

		stmt, err := db.Prepare(qry)
		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}
		defer stmt.Close()

		var rows *sql.Rows
		if len(args) > 0 {
			rows, err = stmt.Query(args...)
		} else {
			rows, err = stmt.Query()
		}

		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}

		for rows.Next() {
			ch <- factory(rows)
		}

		close(ch)
	}()

	return ch
}

// QueryRow gets the query for single result from the database
func (d *Database) QueryRow(qry string, factory func(*sql.Row) interface{}, args ...interface{}) <-chan interface{} {

	ch := make(chan interface{})

	go func() {
		db, err := sql.Open(d.driver, d.connection)
		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}
		defer db.Close()

		stmt, err := db.Prepare(qry)
		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}
		defer stmt.Close()

		var rows *sql.Row
		if len(args) > 0 {
			rows = stmt.QueryRow(args...)
		} else {
			rows = stmt.QueryRow()
		}

		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}

		ch <- factory(rows)

		close(ch)
	}()

	return ch

}

// Execute executes insert, update or delete statements from the database
func (d *Database) Execute(qry string, args ...interface{}) <-chan sql.Result {

	ch := make(chan sql.Result)

	go func() {
		db, err := sql.Open(d.driver, d.connection)
		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}
		defer db.Close()

		stmt, err := db.Prepare(qry)
		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}
		defer stmt.Close()

		r, err := stmt.Exec()
		if err != nil {
			log.Fatalf("ERR Database.Query [%v] %v ", qry, err.Error())
		}

		ch <- r
	}()

	return ch
}
