package main

import (
    "database/sql"
    "log/slog"
    "flag"
    "net/http"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

type application struct {
    logger *slog.Logger
}


func main() {

    addr := flag.String("addr", ":4000", "HTTP network address")

    dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "mySQL data source name")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    db, err := openDB(*dsn)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }

    defer db.Close()

    app := &application{
        logger: logger,
    }

    // register file server as handler for all URL paths which start with
    // "static". Strip /static from matching paths

    logger.Info("starting server", slog.Any("addr", *addr))

    err = http.ListenAndServe(*addr, app.routes())

    logger.Error(err.Error())
    os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
