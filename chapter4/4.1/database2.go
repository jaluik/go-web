package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Uid   int
	Name  string
	Phone string
}

func main() {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/chapter4")
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	rows, err := db.Query("select uid, name, phone from user")
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	u := User{}
	for rows.Next() {
		err := rows.Scan(&u.Uid, &u.Name, &u.Phone)
		if err != nil {
			return
		}
		fmt.Printf("uid: %d, name: %s, phone: %s\n", u.Uid, u.Name, u.Phone)

	}

}
