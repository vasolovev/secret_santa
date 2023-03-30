// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/vasolovev/secret_santa/config"
	v1 "github.com/vasolovev/secret_santa/internal/controller/http/v1"
	"github.com/vasolovev/secret_santa/internal/entity"
	"github.com/vasolovev/secret_santa/internal/repo"
	"github.com/vasolovev/secret_santa/internal/usecase"
	"github.com/vasolovev/secret_santa/pkg/httpserver"
	"github.com/vasolovev/secret_santa/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	var err error
	l := logger.New(cfg.Log.Level)
	var groups []entity.Group
	var participants []entity.Participant

	repos := repo.NewRepositories(groups, participants)
	usecases := usecase.NewUsecases(usecase.Deps{
		Repos: repos,
	})

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, usecases)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())

		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}

	}
}
