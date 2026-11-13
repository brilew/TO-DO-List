package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"TO-DO-List/internal/config"
	"TO-DO-List/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg.Port, cfg.WebDir)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
