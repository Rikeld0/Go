package service

import (
	"database/sql"
	"fmt"

	"../log"
	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var Db Idb = &idb{}

type idb struct {
	Idb
}

var db *sql.DB

//Init ...
func Init() {
	var err error
	db, err = sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		log.Fatal("Error open database.")
	}
}

//Close ...
func Close() {
	db.Close()
}

//Get ...
func (*idb) Get(url string) string {
	rows, err := db.Query("SELECT status FROM resources WHERE url='" + url + "'")
	if err != nil {
		log.Warn("Error request query in database.")
	}
	var statusres string
	statusres = "404 FATAL"
	for rows.Next() {
		rows.Scan(&statusres)
		fmt.Printf("%s\n", statusres)
	}
	return statusres
}

func (*idb) GetAll() []string {
	var list []string
	rows, err := db.Query("SELECT url FROM resources")
	if err != nil {
		log.Warn("Error requset query in database")
	} else {
		for rows.Next() {
			var buf string
			rows.Scan(&buf)
			list = append(list, buf)
		}
	}
	return list
}

//Put ...
func (*idb) Put(url, status string) bool {
	table := `CREATE TABLE IF NOT EXISTS resources (
		id INTEGER PRIMARY KEY,
		url TEXT,
		status TEXT)`
	st, err := db.Prepare(table)
	if err != nil {
		log.Warn("Error request in database.")
		return false
	}
	_, err = st.Exec()
	if err != nil {
		log.Warn("Error request in database")
		return false
	}
	st, err = db.Prepare("INSERT INTO resources (url, status) VALUES (?, ?)")
	if err != nil {
		log.Warn("Error request in database.")
		return false
	}
	_, err = st.Exec(url, status)
	if err != nil {
		log.Warn("Error request in database.")
		return false
	}
	return true
}
