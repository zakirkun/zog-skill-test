package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerContext struct {
	Handler http.Handler
	Host    string

	CertFile interface{}
	KeyFile  interface{}

	Timeout      time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewServer(s ServerContext) ServerContext {
	return ServerContext{
		Host:         s.Host,
		CertFile:     s.CertFile,
		KeyFile:      s.KeyFile,
		Timeout:      s.Timeout,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
	}
}

func (s ServerContext) Run() {
	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer shutdownCancel()

	// Define server options
	server := &http.Server{
		Addr:         s.Host,
		Handler:      s.Handler,
		ReadTimeout:  s.Timeout * time.Second,
		WriteTimeout: s.WriteTimeout * time.Second,
		IdleTimeout:  s.IdleTimeout * time.Second,
	}

	// info
	log.Printf("Server Running on : %v", s.Host)

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTERM)

	// Run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	log.Fatalf("Server is shutting down due to %+v\n", interrupt)

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
