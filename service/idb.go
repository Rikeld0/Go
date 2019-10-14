package service

import (
	"database/sql"
	"fmt"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//Init ...
func Init() {
	db, _ = sql.Open("sqlite3", "./mydb.db")
}

//Close ...
func Close() {
	db.Close()
}

//Get ...
func Get(url string) string {
	rows, _ := db.Query("SELECT status FROM resources WHERE url='" + url + "'")
	var statusres string
	statusres = "404 FATAL"
	for rows.Next() {
		rows.Scan(&statusres)
		fmt.Printf("%s\n", statusres)
	}
	return statusres
}

//Put ...
func Put(url, status string) bool {
	table := `CREATE TABLE IF NOT EXISTS resources (
		id INTEGER PRIMARY KEY,
		url TEXT,
		status TEXT)`
	st, err := db.Prepare(table)
	if err != nil {
		return false
	}
	_, err = st.Exec()
	if err != nil {
		return false
	}
	st, err = db.Prepare("INSERT INTO resources (url, status) VALUES (?, ?)")
	if err != nil {
		return false
	}
	_, err = st.Exec(url, status)
	if err != nil {
		return false
	}
	return true
}
