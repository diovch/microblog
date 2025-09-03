package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/diovch/microblog/config"
	"github.com/diovch/microblog/internal/handlers"
	"github.com/diovch/microblog/internal/logger"
	"github.com/diovch/microblog/internal/repo"
	"github.com/diovch/microblog/internal/service"
	"github.com/gorilla/mux"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(100)
	defer l.Close()

	memDb := repo.NewMemoryRepo()
	wp := service.NewWorkerPool()
	defer wp.Close()

	userHandler := handlers.NewUserHandler(memDb, l)
	postHandler := handlers.NewPostHandler(memDb, wp, l)

	r := mux.NewRouter()
	r.HandleFunc("/register", userHandler.RegisterHandler).Methods(http.MethodPost)
	r.HandleFunc("/posts", postHandler.CreatePostHandler).Methods(http.MethodPost)
	r.HandleFunc("/posts", postHandler.GetFeedHandler).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}/like", postHandler.LikePostHandler).Methods(http.MethodGet)

	// FAQ: Should httpServer be in other package?
	srv := &http.Server{
		Addr:    ":" + string(rune(cfg.HTTP.Port)),
		Handler: r,
	}

	go func() {
		l.LogInfo(fmt.Sprint("Starting server on port ", cfg.HTTP.Port))
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Shutdown

	if err := srv.Shutdown(ctx); err != nil {
		l.LogError(fmt.Sprint("app - Run - httpServer.Shutdown: %w", err))
	}
	l.LogInfo("Shutting down gracefully")
}
