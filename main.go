// Copyright 2014-2017 Peter Hellberg.
// Released under the terms of the MIT license.

// wiki is a tiny wiki using BoltDB and Blackfriday.
//
// Installation
//
// You can install wiki with go get:
//
//     go get -u github.com/peterhellberg/wiki
//
// Usage
//
// You can specify one (optional) parameter -db
//
//     PORT=7272 wiki -db="/tmp/foo.db"
//
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/peterhellberg/wiki/server"
)

var basePath string

func main() {
	flag.StringVar(&basePath, "files", "/tmp/", "Path to wiki files")

	flag.Parse()

	// Setup the logger used by the server
	logger := log.New(os.Stdout, "", 0)

	hs := setup(logger, basePath)

	go graceful(hs, 8*time.Second)

	logger.Println("Listening on http://" + hs.Addr)
	if err := hs.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func setup(logger *log.Logger, basePath string) *http.Server {
	return &http.Server{
		Addr:         addr(),
		Handler:      server.New(logger, basePath),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}
}

func addr() string {
	if port := os.Getenv("PORT"); port != "" {
		return "127.0.0.1:" + port
	}

	return "127.0.0.1:7272"
}

func graceful(hs *http.Server, timeout time.Duration) {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	hs.Shutdown(ctx)
}
