package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql" 
)

type MySQLConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func NewMySQLConnection(cfg MySQLConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("erro ao conectar no mysql: %w", err)
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(10)
    db.SetConnMaxLifetime(5 * time.Minute)

    log.Println("mysql conectado com sucesso")
    return db, nil
}