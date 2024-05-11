package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	restserver "github.com/rarrazaan/be-player-performance-app/cmd"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"github.com/rarrazaan/be-player-performance-app/internal/dependency"
)

var (
	shutdownTimeout = 5 * time.Second
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := config.GetConfig()
	dependency := dependency.NewDependency(ctx, config)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	cleanUpChan := make(chan error)

	go func() {
		log.Printf("receiving signal %s, shutting down...\n", <-quitChan)
		cancel()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		select {
		case <-dependency.Close(shutdownCtx):
			cleanUpChan <- nil
		case <-time.After(shutdownTimeout * 3 / 2):
			cleanUpChan <- errors.New("graceful shutdown timeout")
		}
	}()

	restserver.Serve(dependency)

	if err := <-cleanUpChan; err != nil {
		log.Println("graceful shutdown timeout, force shutdown...")
		os.Exit(1)
	}
	log.Println("app shutdown gracefully")
}
