package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("sqlite", "cotacoes.db")
	if err != nil {
		panic(fmt.Errorf("erro ao abrir conexão com banco: %w", err))
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Verificar conexão
	if err = db.Ping(); err != nil {
		panic(fmt.Errorf("erro ao conectar com banco: %w", err))
	}

	database := NewDatabase(db)
	err = database.Migrations()
	if err != nil {
		panic(err)
	}

	api := NewAPI(database)
	api.StartServer()

}
