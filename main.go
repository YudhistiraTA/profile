package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YudhistiraTA/profile/db"
	"github.com/YudhistiraTA/profile/services"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

func main() {
	// env init
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// context init
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	defer cancel()

	// db init
	db, err := db.NewDatabase(ctx)
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}
	// db ping
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("db ping failed: %v", err)
	}
	defer db.Close(ctx)

	// server init
	addr := os.Getenv("ADDRESS")
	srv := services.NewServer(db, addr)
	rungroup, ctx := errgroup.WithContext(ctx)
	rungroup.Go(func() error {
		if er := srv.ListenAndServe(); er != nil && !errors.Is(er, http.ErrServerClosed) {
			return fmt.Errorf("listen and server %w", er)
		}
		return nil
	})
	rungroup.Go(func() error {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if er := srv.Shutdown(shutdownCtx); er != nil {
			return fmt.Errorf("shutdown http server %w", er)
		}

		return nil
	})
	if err := rungroup.Wait(); err != nil {
		log.Fatal(fmt.Errorf("run group exited because of error: %v", err))
		return
	}
	fmt.Print("server exited properly")
}
