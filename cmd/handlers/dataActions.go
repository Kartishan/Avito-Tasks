package handlers

import (
	"Avito-Tasks/cmd/data"
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

// Подключение к БД
func ConnectDB() {
	var cfg Сonfig

	flag.IntVar(&cfg.Port, "port", 4000, "Api server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.Db.Dsn, "db-dsn",
		"postgres://postgres:root@0.0.0.0:5432/AVITO-TASKS?sslmode=disable", "PostgeSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Printf("database connection pool established")

	app := &Application{
		Config: cfg,
		Logger: logger,
		Models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.Env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

// открытие БД
func openDB(cfg Сonfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
