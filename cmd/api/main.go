package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davidlick/supermarket-api/internal/http"
	"github.com/davidlick/supermarket-api/internal/produce"
	"github.com/davidlick/supermarket-api/pkg/ramdb"
	"github.com/sirupsen/logrus"
)

var (
	cfg    config
	logger *logrus.Logger
	err    error
)

func init() {
	cfg, err = load()
	if err != nil {
		log.Fatal(err)
	}

	ll, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger = logrus.New()
	logger.SetLevel(ll)
}

func main() {
	db := ramdb.NewDatabase()
	err = db.CreateTable("produce", produce.KeyProduceCode)
	if err != nil {
		logger.Fatal(err)
	}

	produceSvc := produce.NewService(db.From("produce"))

	server := http.NewServer(cfg.APIPort, logger, cfg.Env, produceSvc)

	// Allow app to listen for OS Interrupts and SIGTERMS.
	serverErrors := make(chan error, 1)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	go func() {
		serverErrors <- server.Run()
	}()

	// Handling for server errors and OS signals.
	select {
	case err := <-serverErrors:
		log.Fatal("error starting server: %w", err.Error())
	case <-osSignals:
		log.Println("starting server shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatalf("error shutting down http server: %v", err.Error())
		}
	}
}
