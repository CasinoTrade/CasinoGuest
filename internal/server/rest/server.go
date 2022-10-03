package rest

import (
	"context"
	"net/http"

	"github.com/CasinoTrade/CasinoGuest/internal/model/config"
	"github.com/CasinoTrade/CasinoGuest/internal/model/log"
	model "github.com/CasinoTrade/CasinoGuest/internal/model/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CasinoREST struct {
	ctx    context.Context
	cancel context.CancelFunc

	cfg config.Server
	log log.Logger
	e   *echo.Echo
}

func (s *CasinoREST) Start() {
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

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Number int }{Number: 42})
	})

	e.Logger.Fatal(e.Start(s.cfg.Address))
}

func (s *CasinoREST) Stop() {
	if s.cancel != nil {
		defer s.cancel()
	}
	// NB: Shutdown waits for active connections, without interrupting them.
	// May be we should consider using Stop.
	s.e.Shutdown(context.Background())
}

func New(cfg config.Server, log log.Logger) model.Casino {
	ctx, cancel := context.WithCancel(context.Background())
	return &CasinoREST{
		cfg:    cfg,
		log:    log,
		ctx:    ctx,
		cancel: cancel,
	}
}
