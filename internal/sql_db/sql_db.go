package sql_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("db_host"), os.Getenv("db_port"), os.Getenv("db_user"), os.Getenv("db_name"), os.Getenv("db_password"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func InsertTerminStatus(db *sql.DB, status bool) {
	sqlStatement := `
	INSERT INTO terminStatus (status) 
	VALUES ($1)
	RETURNING id, created_at`
	var id int
	var createdAt string

	err := db.QueryRow(sqlStatement, status).Scan(&id, &createdAt)
	if err != nil {
		log.Fatal("Error inserting data: ", err)
	}

}
