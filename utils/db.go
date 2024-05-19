package utils

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/ClickHouse/clickhouse-go"
)

func NewDB() (*sqlx.DB, error) {
    db, err := sqlx.Open("clickhouse", "tcp://localhost:9000?debug=true")
    if err != nil {
        return nil, err
    }
    return db, nil
}
