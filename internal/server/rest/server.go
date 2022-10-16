package rest

import (
	"context"
	"net/http"

	"github.com/CasinoTrade/CasinoGuest/internal/model/config"
	"github.com/CasinoTrade/CasinoGuest/internal/model/log"
	model "github.com/CasinoTrade/CasinoGuest/internal/model/server"
	"github.com/CasinoTrade/CasinoGuest/internal/server"

	"github.com/gofiber/fiber/v2"
	fibrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

type CasinoREST struct {
	ctx    context.Context
	cancel context.CancelFunc

	cfg config.Server
	log log.Logger
	e   *fiber.App

	base *server.Casino
}

func (s *CasinoREST) Start() {
	s.log.Debug("Preparing REST server")
	e := fiber.New()
	s.e = e

	logger := s.log.WithSource("rest-middleware")
	e.Use(func(c *fiber.Ctx) error {
		logger.WithFields(
			log.Field{"method", c.Method()},
			log.Field{"ip", c.IP()},
			log.Field{"path", c.OriginalURL()},
		).Info("Request handled")
		return c.Next()
	})
	e.Use(fibrecover.New())

	handler := newHandler(s.base)

	e.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hellow!")
	})

	e.Get("/ping", handler.Ping)

	go s.run()
}

func (s *CasinoREST) run() {
	log := s.log.WithSource("server-runner")
	log.Infof("Starting server with config: %s", s.cfg)
	if err := s.e.Listen(s.cfg.Address); err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
	log.Info("Server closed")
}

func (s *CasinoREST) Stop() {
	s.log.Info("Shutdown server in progress")
	if err := s.e.Shutdown(); err != nil {
		s.log.Fatalf("Shutdown error: %v", err)
	}

	if s.cancel != nil {
		s.cancel()
	}
}

func New(cfg config.Server, log log.Logger, base *server.Casino) model.Casino {
	ctx, cancel := context.WithCancel(context.Background())
	return &CasinoREST{
		cfg:    cfg,
		log:    log,
		ctx:    ctx,
		cancel: cancel,
		base:   base,
	}
}
