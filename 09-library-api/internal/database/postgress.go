package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gizaw/09-library-api/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(cfg *config.Config )(*sql.DB, error) {
	db,err := sql.Open("pgx",cfg.DatabaseURL )

	if err != nil {
		return nil, fmt.Errorf("open database : %w", err)
	}

	cxt, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err := db.PingContext(cxt); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db : %w", err)
	}

	// connection poll config 
	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(cfg.DBConnMaxIdleTime) * time.Minute)

	return db, nil 
}

