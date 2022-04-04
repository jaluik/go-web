package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/chapter4")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init database failed: %v \n", err)
		return
	}
	smst, err := db.Prepare("insert  into user (uid, name, phone) values  (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(smst *sql.Stmt) {
		err := smst.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(smst)
	_, err = smst.Exec(1, "jaluik", "18580001234")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = smst.Exec(2, "jack", "18580004321")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("insert success")

}
