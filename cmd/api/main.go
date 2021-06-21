package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
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
	initProduce(produceSvc)

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

func initProduce(produceSvc http.ProduceService) {
	if cfg.DMLInitFile == "" {
		logger.Info("no init file provided, database will be empty")
		return
	}

	f, err := os.Open(cfg.DMLInitFile)
	if err != nil {
		logger.Fatalf("could not open file: %s", cfg.DMLInitFile)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Fatal("could not read file: %s", cfg.DMLInitFile)
	}

	var items []produce.Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		logger.Fatalf("could not unmarshal items: %v", err)
	}

	err = produceSvc.Add(items)
	if err != nil {
		logger.Fatal("failed to add items to produce service")
	}
}
