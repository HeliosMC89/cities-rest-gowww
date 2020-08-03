package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gowww/router"
	"github.com/heliosmc89/api-rest-gowww/api/repository"
	"github.com/heliosmc89/api-rest-gowww/api/routes"
	"github.com/heliosmc89/api-rest-gowww/config"
	"github.com/jmoiron/sqlx"
)

// EnvName is the environment Name int representation
// Using iota, 1 (Production) is the lowest,
// 2 (Staging) is 2nd lowest, and so on...
type EnvName uint8

// EnvName of environment.
const (
	Production  EnvName = iota + 1 // Production (1)
	Staging                        // Staging (2)
	Testing                        // Testing (3)
	Development                    // Development (4)
)

type Server struct {
	ENVName EnvName
	DB      *sqlx.DB
	Router  *router.Router
	Logger  *log.Logger
	Config  *config.Config
	Version string
}

func (s *Server) loadConfig() {
	s.Config = config.GetConfig()
}

func (s *Server) setVersion() {
	s.Version = "1.0.0"
}

func (s *Server) logger() {
	s.Logger = log.New(os.Stdout, "api/", log.LstdFlags|log.Lshortfile)
}

func (s *Server) loadDatabase() {
	uri := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		s.Config.DB.Username,
		s.Config.DB.Password,
		s.Config.DB.Name,
		s.Config.DB.Host,
		s.Config.DB.Port,
		s.Config.DB.SslMode)
	var err error
	s.DB, err = repository.NewDB(s.Config.DB.Dialect, uri)
	if err != nil {
		s.Logger.Panicln(err)
	}

}

func (s *Server) loadRoutes() {
	s.Router = routes.NewRouter(s.Logger, s.DB)
}

func NewServer() *Server {
	server := &Server{}
	server.ENVName = Development
	server.logger()
	server.setVersion()
	server.loadConfig()
	server.loadDatabase()
	server.loadRoutes()
	return server
}

func (s *Server) Run(addr string) {
	srv := http.Server{
		Addr:    addr,
		Handler: s.Router,

		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shutdown.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Http server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	s.Logger.Printf("Server starting on port %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		// Error stating or closing listener:
		s.Logger.Printf("Http server Listen And Serve: %v", err)
	}
	<-idleConnsClosed
}
