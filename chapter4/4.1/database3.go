package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/chapter4")
	if err != nil {
		log.Fatal(err)
		return
	}
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
		fmt.Printf("begin trans failed, err:%v \n", err)
		return
	}
	_, err = tx.Exec("update user set name='james' where uid = ?", 2)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		fmt.Printf("exec sql1 failed error: %v", err)
		return
	}
	_, err = tx.Exec("update user set name='james' where uid = ?", 3)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		fmt.Printf("exec sql2 failed error: %v", err)
		return

	}
	err = tx.Commit()
	if err != nil {
		return
	}
	fmt.Println("exec transaction success!")
}
