package main

import (

	"github.com/jmoiron/sqlx"
	"github.com/ussg43/opselling/internal/store"
)

type config struct {
	dbConn   string
	httpPort string
	rpcPort  string
}

type application struct {
	config  config
	db      *sqlx.DB
	storage store.Storage
}



func (a *application) mount() {
	
}
