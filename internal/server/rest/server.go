package rest

import (
	"context"

	"github.com/CasinoTrade/CasinoGuest/internal/model/config"
	"github.com/CasinoTrade/CasinoGuest/internal/model/log"
	model "github.com/CasinoTrade/CasinoGuest/internal/model/server"

	"github.com/gofiber/fiber/v2"
	fibrecover "github.com/gofiber/fiber/v2/middleware/recover"
)

type CasinoREST struct {
	ctx    context.Context
	cancel context.CancelFunc

	cfg config.Server
	log log.Logger
	e   *fiber.App

	base model.Base
}

func New(cfg config.Server, log log.Logger, base model.Base) model.Casino {
	ctx, cancel := context.WithCancel(context.Background())
	return &CasinoREST{
		cfg:    cfg,
		log:    log,
		ctx:    ctx,
		cancel: cancel,
		base:   base,
	}
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

func (s *CasinoREST) Stop() {
	s.log.Info("Server shutdown in progress")
	if err := s.e.Shutdown(); err != nil {
		s.log.Fatalf("Shutdown error: %v", err)
	}

	if s.cancel != nil {
		s.cancel()
	}
	s.log.Debug("Sever shutdown done")
}

func (s *CasinoREST) run() {
	log := s.log.WithSource("rest-runner")
	log.Debugf("Starting server with config: %#v", s.cfg)
	if err := s.e.Listen(s.cfg.Address); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	log.Info("Server closed")
}
