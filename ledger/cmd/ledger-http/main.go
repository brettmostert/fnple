package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"os"

	"github.com/brettmostert/fnple/ledger/internal/common"
	api "github.com/brettmostert/fnple/ledger/server/http"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// todo - server - http
	connString := os.Getenv("DATABASE_URL") + "/ledger"

	conn, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// this needs to be at the http server level on shutdown
	defer conn.Close()

	ctx := &common.AppContext{
		Db: conn,
	}

	api := api.NewApi(ctx)
	// todo - add to config
	server := http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      api,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	server.ListenAndServe()
}
