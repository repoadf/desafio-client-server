package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

func (d *Database) Migrations() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT NOT NULL,
		criado_em DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := d.db.ExecContext(context.Background(), createTableSQL)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela: %w", err)
	}
	return nil
}

func (d *Database) SaveCotacao(bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	insertSQL := "INSERT INTO cotacoes (bid) VALUES (?)"
	_, err := d.db.ExecContext(ctx, insertSQL, bid)
	if err != nil {
		return fmt.Errorf("erro ao inserir cotação: %w", err)
	}
	return nil
}
