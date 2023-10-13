package main

import (
    "log/slog"
    "flag"
    "net/http"
    "os"
)

type application struct {
    logger *slog.Logger
}


func main() {

    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    app := &application{
        logger: logger,
    }

    // register file server as handler for all URL paths which start with
    // "static". Strip /static from matching paths

    logger.Info("starting server", slog.Any("addr", *addr))

    err := http.ListenAndServe(*addr, app.routes())

    logger.Error(err.Error())
    os.Exit(1)
}
