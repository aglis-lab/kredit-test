package main

import (
	"context"
	"errors"
	"fmt"
	"kreditplus/src/app"
	"kreditplus/src/middleware"
	"kreditplus/src/tracer"
	v1 "kreditplus/src/v1"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Init app context
	if err := app.Init(ctx); err != nil {
		log.Panic(err)
	}

	// Setup Otel SDK
	otelShutdown, err := tracer.SetupOTelSDK(ctx, app.Config().OltpGRPCProvider, app.Config().ServiceName, app.Config().ServiceVersion)
	if err != nil {
		log.Panic(err)
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
		log.Println(err)
	}()

	// Init Router
	address := fmt.Sprintf(":%d", app.Config().BindAddress)
	r := initRouter(ctx, address)

	// Start HTTP Server
	srv := &http.Server{
		Addr:         address,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}
	srvErr := make(chan error, 1)

	log.Printf("Starting service on %s", address)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
}

func initRouter(ctx context.Context, address string) *chi.Mux {
	r := chi.NewRouter()
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.RequestIDContext(middleware.DefaultGenerator))
	r.Use(middleware.RequestAttributesContext)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	deps := v1.Dependencies(ctx)

	r.Route("/v1", func(r chi.Router) {
		v1.Router(r, deps)
	})

	return r
}
