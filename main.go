package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "mypassword",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "projectmanager",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	server := NewAPIServer(":8080", store)
	server.Serve()
}
