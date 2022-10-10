package rest

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/CasinoTrade/CasinoGuest/internal/model/config"
	"github.com/CasinoTrade/CasinoGuest/internal/model/log"
	model "github.com/CasinoTrade/CasinoGuest/internal/model/server"
	"github.com/CasinoTrade/CasinoGuest/internal/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const shutdownTimeout = 2 * time.Minute

type CasinoREST struct {
	ctx    context.Context
	cancel context.CancelFunc

	cfg config.Server
	log log.Logger
	e   *echo.Echo

	base *server.Casino
}

func (s *CasinoREST) Start() {
	s.log.Debug("Preparing REST server")
	e := echo.New()
	e.HideBanner = true
	s.e = e

	logger := s.log.WithSource("rest-middleware")
	lc := middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogMethod:   true,
		LogURIPath:  true,
		LogError:    true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			e := v.Error
			if e == nil {
				logger.WithFields(
					log.Field{"method", v.Method},
					log.Field{"ip", v.RemoteIP},
					log.Field{"path", v.URIPath},
				).Info("Request handled")
				return nil
			}
			logger.WithFields(
				log.Field{"method", v.Method},
				log.Field{"ip", v.RemoteIP},
				log.Field{"path", v.URIPath},
				log.Field{"err", v.Error.Error()},
			).Warn("Request failed")
			return e
		},
	}

	e.Use(middleware.RequestLoggerWithConfig(lc))
	e.Use(middleware.Recover())

	handler := newHandler(s.base)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello!")
	})

	e.GET("/ping", handler.Ping)

	// TODO: Check, we still allowed to log evryting needed or provide logger for echo
	e.Logger.SetOutput(ioutil.Discard)

	s.run()
}

func (s *CasinoREST) run() {
	log := s.log.WithSource("server-runner")
	log.Infof("Starting server with config: %s", s.cfg)
	if err := s.e.Start(s.cfg.Address); err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
	log.Info("Server closed")
}

func (s *CasinoREST) Stop() {
	s.log.Info("Shutdown server in progress")
	// NB: Shutdown waits for active connections, without interrupting them.
	ctx, cancel := context.WithTimeout(s.ctx, shutdownTimeout)
	if err := s.e.Shutdown(ctx); err != nil {
		s.log.Fatalf("Shutdown error: %v", err)
	}

	cancel() // not realy needed, juts dont disturb linters
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
