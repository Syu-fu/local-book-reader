package infra

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	err := godotenv.Load("./.env")

	if err != nil {
		fmt.Printf("Can't read .env file: %v", err)
	}
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	port := os.Getenv("MYSQL_PORT")
	conn, err := sql.Open("mysql", user+":"+pass+"@tcp(mysql:"+port+")/local_book_reader")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
