package main

import (
	"database/sql"
	"log"

	"github.com/skoal2007/golang-simple-bank/api"
	db "github.com/skoal2007/golang-simple-bank/db/sqlc"
	"github.com/skoal2007/golang-simple-bank/util"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	// dbSource = "postgresql://root:secret@192.168.0.139:5432/simple_bank?sslmode=disable"
// 	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot read config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
