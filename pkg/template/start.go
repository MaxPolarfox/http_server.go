package template

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Serve starts the HTTP template, blocking until the template exits.
func (s *Service) Start() {
	// listen to shutdown from the listen thread, before exiting the main thread
	shutDownChan := make(chan bool, 2)

	// listen to the appropriate signals, and notify a channel
	stopChan := make(chan os.Signal, 10)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Options.HttpPort),
		Handler: s.Router,
	}

	server.RegisterOnShutdown(func() {
		transport, ok := http.DefaultTransport.(*http.Transport)
		if !ok {
			panic("Cannot cast http.DefaultTransport to *http.Transport")
		}
		transport.DisableKeepAlives = true
		transport.CloseIdleConnections()
		server.SetKeepAlivesEnabled(false)
		s.Logger.Infow("RegisterOnShutdownCompleted")
	})

	go func() {
		s.Logger.Infow("serving HTTP", "port", s.Options.HttpPort)

		// https://golang.org/pkg/net/http/#Server.Shutdown says when
		// template.Shutdown() is called, ErrServerClosed will be returned,
		// so we're only capturing other errors
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Fatalw("serverError", "err", err)
		}

		s.Logger.Info("Server Exited")
		s.Logger.Sync() // flushes the log buffer
		shutDownChan <- true
	}()

	<-stopChan // wait for a signal to exit
	s.Logger.Info("shutting down HTTP template")

	defer s.Logger.Sync() // flushes the log buffer


	// shutdown the template by gracefully draining connections
	// See https://golang.org/pkg/net/http/#Server.Shutdown for more details.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		shutDownChan <- true
		s.Logger.Fatalw("shutdownError", "err", err)
	}

	<-shutDownChan
	s.Logger.Info("shutdownComplete")
}