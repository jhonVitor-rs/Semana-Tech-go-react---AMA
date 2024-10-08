package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jhonVitor-rs/Semana-Tech-go-react---AMA/go-server/internal/api"
	"github.com/jhonVitor-rs/Semana-Tech-go-react---AMA/go-server/internal/store/pgstore"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != err {
		slog.Warn("Failed to get environments variables")
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("WSRS_DATABASE_USER"),
		os.Getenv("WSRS_DATABASE_PASSWORD"),
		os.Getenv("WSRS_DATABASE_HOST"),
		os.Getenv("WSRS_DATABASE_PORT"),
		os.Getenv("WSRS_DATABASE_NAME"),
	))
	if err != nil {
		slog.Warn("Failed to create pool")
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Warn("Failed to pool.Ping!")
		panic(err)
	}

	handler := api.NewHandler(pgstore.New(pool))

	go func() {
		if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				slog.Warn("Failed to get environments variables")
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
