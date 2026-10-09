package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"ssp-sp-messaging/api/internal/adapters/handler"
	"ssp-sp-messaging/api/internal/adapters/queue"
	"ssp-sp-messaging/api/internal/config"
	"ssp-sp-messaging/api/internal/core/services"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()

	publisher, err := queue.NewPublisher(cfg.RabbitMQURL, cfg.Queue)
	if err != nil {
		log.Fatalf("rabbitmq: %v", err)
	}
	defer publisher.Close()

	// core
	svc := services.NewMessageService(publisher)

	// inbound adapter
	srv := &http.Server{Addr: cfg.HTTPAddr, Handler: handler.NewRouter(svc)}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("listening on %s", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}
