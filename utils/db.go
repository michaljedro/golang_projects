package utils

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/kshvakov/clickhouse"
    "log"
)

var DB *sqlx.DB

func InitDB() {
    var err error
    DB, err = sqlx.Open("clickhouse", "tcp://localhost:9000?username=default&password=&database=mydb")
    if err != nil {
        log.Fatalln(err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalln(err)
    }

    log.Println("Connected to ClickHouse!")
}
