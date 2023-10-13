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

    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))

    // register file server as handler for all URL paths which start with
    // "static". Strip /static from matching paths
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/view", app.snippetView)
    mux.HandleFunc("/snippet/create", app.snippetCreate)

    logger.Info("starting server", slog.Any("addr", *addr))

    err := http.ListenAndServe(*addr, mux)

    logger.Error(err.Error())
    os.Exit(1)
}
