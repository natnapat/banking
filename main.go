package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/natnapat/banking/api"
	db "github.com/natnapat/banking/db/sqlc"
	"github.com/natnapat/banking/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
