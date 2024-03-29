package dbports

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectToDataBase() *sql.DB {

	db_user := os.Getenv("db_user")

	db_passwd := os.Getenv("db_passwd")

	db_addr := os.Getenv("db_addr")

	db_db := os.Getenv("db_db")

	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_passwd, db_addr, db_db)

	var err error

	db, err := sql.Open("mysql", s)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Connecting to database...")

	if err := db.Ping(); err != nil {

		log.Fatalln(err)

	}

	fmt.Println("Connnected to the database")

	return db

}
