package mysql

import (
	"database/sql"
	"fmt"
	"gin-react-auth/utils/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	config.LoadEnv()
	var (
		USER     = os.Getenv("DB_USER")
		PASSWORD = os.Getenv("DB_PASSWORD")
		HOST     = os.Getenv("DB_HOST")
		PORT     = os.Getenv("DB_PORT")
		DB_NAME  = os.Getenv("DB_NAME")
	)
	// username:password@tcp(host:port)/user_schema
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", USER, PASSWORD, HOST, PORT, DB_NAME)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully configured!")
}
